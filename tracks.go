package main

import (
	"github.com/astaxie/beego"
)

type (
	TracksController struct {
		beego.Controller
	}
)

// TopTracks handles the /top_tracks endpoint,
// returns a list of the user's top tracks
func (c *TracksController) TopTracks() {
	token := c.Ctx.Input.Header("Token")
	topTracks, err := tracksService.TopTracks(token)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Data["json"] = topTracks
	c.ServeJSON()
}
