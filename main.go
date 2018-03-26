package main

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/nbatchelder/soundpals-be/tracks"
)

var (
	tracksService tracks.Service
)

func main() {
	httpClient := http.Client{}

	tracksService = tracks.NewService(&httpClient)

	// routing
	beego.Router("/top_tracks", &TracksController{}, "get:TopTracks")
	beego.Run()
}
