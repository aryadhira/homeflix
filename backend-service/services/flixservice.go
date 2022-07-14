package services

import (
	"context"
	"homeflix/engine"
	"homeflix/helper"
	. "homeflix/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FlixService struct{}

func (fs *FlixService) RunService() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(fs.HttpMiddleware)

	router.GET("/getmovies", fs.GetMovieData)
	router.GET("/getmoviebyid/:id", fs.GetMovieById)
	router.GET("/updatemovie", fs.UpdateCollection)
	router.Run("localhost:9090")
}

func (fs *FlixService) HttpMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.Next()
}

func (fs *FlixService) GetMovieData(c *gin.Context) {
	ctx := context.Background()
	db, err := helper.DBConnect()
	if err != nil {
		log.Println(err.Error())
	}
	opts := options.Find().SetSort(bson.M{"dateuploaded": -1})
	csr, err := db.Collection("movies").Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Println(err.Error())
	}

	defer csr.Close(ctx)

	results := []Movie{}
	for csr.Next(ctx) {
		row := Movie{}
		err := csr.Decode(&row)
		if err != nil {
			log.Println(err.Error())
		}
		results = append(results, row)
	}

	res := helper.SetHttpResponse("", results)

	c.JSON(http.StatusOK, res)
}

func (fs *FlixService) GetMovieById(c *gin.Context) {
	ctx := context.Background()
	db, err := helper.DBConnect()
	if err != nil {
		log.Println(err.Error())
	}

	id := c.Param("id")
	idnum, _ := strconv.Atoi(id)

	csr, err := db.Collection("movies").Find(ctx, bson.M{"id": idnum})
	if err != nil {
		log.Println(err.Error())
	}

	defer csr.Close(ctx)

	results := Movie{}
	for csr.Next(ctx) {
		err := csr.Decode(&results)
		if err != nil {
			log.Println(err.Error())
		}
	}

	res := helper.SetHttpResponse("", results)

	c.JSON(http.StatusOK, res)
}

func (fs *FlixService) UpdateCollection(c *gin.Context) {
	res := "Update Data Success"

	grabber := new(engine.MovieEngine)
	err := grabber.UpdateMovieCollection()
	if err != nil {
		res = err.Error()
	}

	result := helper.SetHttpResponse(res, nil)

	c.JSON(http.StatusOK, result)
}
