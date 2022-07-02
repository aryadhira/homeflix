package main

import (
	"homeflix/engine"
)

func main() {
	// engine := new(engine.MovieEngine)
	// err := engine.SaveMovieData()
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Println("Data Movie Saved")
	testengine := new(engine.TestEngine)
	testengine.GetMovieList()
}
