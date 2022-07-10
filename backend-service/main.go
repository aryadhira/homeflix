package main

import (
	"homeflix/services"
)

func main() {
	// engine := new(engine.MovieEngine)
	// err := engine.SaveMovieData()
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Println("Data Movie Saved")
	// testengine := new(engine.TestEngine)
	// testengine.GetMovieList()
	services := services.FlixService{}
	services.RunService()
}
