package engine

import (
	"bufio"
	"context"
	"fmt"
	"homeflix/helper"
	"homeflix/models"
	"log"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type TestEngine struct{}

func (te *TestEngine) GetMovieList() {
	ctx := context.Background()
	db, err := helper.DBConnect()
	if err != nil {
		log.Println(err.Error())
	}

	csr, err := db.Collection("movies").Find(ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
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

	for i, each := range results {
		qualities := each.Torrents
		qualitystr := "|"
		for _, eachtorrent := range qualities {
			qualitystr += eachtorrent.Type + " " + eachtorrent.Quality + "|"
		}
		fmt.Println(i+1, ".", each.TitleLong, "", qualitystr)
	}

	//Select Movie & quality
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose Movie to Play : ")
	no, _ := reader.ReadString('\n')
	strsplit := strings.Split(no, "\r\n")
	idxno, err := strconv.Atoi(strsplit[0])
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Select Quality")
	qual, _ := reader.ReadString('\n')
	strqual := strings.Split(qual, " ")
	_type := strqual[0]
	_qlty := strqual[1]
	movieidx := idxno - 1
	torrentidx := helper.GetTorrentByQuality(results[movieidx].Torrents, _type, _qlty)

	log.Println("Generating Magnet URL...")
	magnet := helper.CreateMagnetUrl(results[movieidx], torrentidx)

	fmt.Println("Playing ", results[idxno-1].TitleEnglish, qual, "...")

	stream := new(StreamerEngine)
	stream.StartStreaming(magnet)
}
