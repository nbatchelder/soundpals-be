package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type (
	UsersController struct {
		beego.Controller
	}

	Artist struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	Album struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	Track struct {
		Id      string   `json:"id"`
		Name    string   `json:"name"`
		Artists []Artist `json:"artists"`
		Album   Album    `json:"album`
	}

	TopTracksResponse struct {
		Tracks []Track `json:"items"`
	}
)

func (c *UsersController) TopTracks() {
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/tracks?time_range=medium_term", nil)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	req.Header.Set("Authorization", "Bearer BQA6vfUr8QplI4Hfd2XrD661Vq6gZAOCwDo5Y4J3pGlG7-BeeKZDbKp5_yXzXg3rpenVlk9pd3S-rYnskE13k9APwOwbq98QEVrDgoNFU9Br5uGYE_8sJBp3VikckdkKwdMmsCPmafWL2AI7S5TSzMpqRxR1bo7xfFugH5SDvg")
	res, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	defer res.Body.Close()

	var topTracks TopTracksResponse
	err = json.NewDecoder(res.Body).Decode(&topTracks)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	fmt.Println("top tracks", topTracks.Tracks)

	c.Data["json"] = topTracks.Tracks
	c.ServeJSON()
}
