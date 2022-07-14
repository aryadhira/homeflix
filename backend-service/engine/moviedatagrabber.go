package engine

import (
	"context"
	"encoding/json"
	"homeflix/configuration"
	"homeflix/helper"
	"homeflix/models"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type MovieEngine struct{}

func (me *MovieEngine) GetMovieData(page int) (models.ApiResult, error) {
	log.Println("Getting Movie Data From YTS...")

	cfg := new(configuration.Config)
	config := cfg.GetConfig()
	limit := config.Limit
	pagenum := strconv.Itoa(page)

	var err error
	var client = &http.Client{}
	var apiurl = config.YtsAPIURL + "?limit=" + limit + "&page=" + pagenum + "&sort_by=year&order_by=desc"
	results := models.ApiResult{}

	request, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		log.Println(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		log.Println(err.Error())
	}

	return results, err
}

func (me *MovieEngine) SaveMovieData(pagenum int) error {
	//Grab Movie Data
	ytsres, err := me.GetMovieData(pagenum)
	moviedata := []models.Movie{}
	if err != nil {
		return err
	}
	moviedata = ytsres.Data.Movies

	//Save Movie Data to Database
	ctx := context.Background()
	db, err := helper.DBConnect()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	csr, err := db.Collection("movies").Find(ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer csr.Close(ctx)

	results := make([]models.Movie, 0)
	for csr.Next(ctx) {
		var row models.Movie
		err := csr.Decode(&row)
		if err != nil {
			log.Println(err.Error())
		}

		results = append(results, row)
	}

	moviemap := make(map[string]string)
	for _, each := range results {
		moviemap[each.Title] = each.TitleEnglish
	}

	if len(moviedata) > 0 {
		log.Println("Saving ", len(moviedata), " moviedata")
		log.Println("Inserting Movie Data...")
		for _, each := range moviedata {
			isExist := me.IsMovieExist(moviemap, each.Title)
			if isExist == false {
				_, err := db.Collection("movies").InsertOne(ctx, each)
				if err != nil {
					return err
				}
			}
		}

	} else {
		log.Println("No Movie Data")
	}

	return nil
}

func (me *MovieEngine) IsMovieExist(moviemap map[string]string, title string) bool {
	isExist := false
	if moviemap[title] != "" {
		isExist = true
	}
	return isExist
}

func (me *MovieEngine) UpdateMovieCollection() error {
	cfg := new(configuration.Config)
	config := cfg.GetConfig()
	page := config.Page

	pagetotal, _ := strconv.Atoi(page)

	for i := 0; i < pagetotal; i++ {
		err := me.SaveMovieData(i + 1)
		if err != nil {
			return err
		}
	}

	return nil
}
