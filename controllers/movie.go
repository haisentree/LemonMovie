package controllers

import (
	"LemonMovie/models"
	"github.com/beego/beego/v2/server/web"
)

type MovieController struct {
	web.Controller
}

func (c *MovieController) GetHome() {
	movieListModel := models.NewMovieModel()
	movieList := movieListModel.FindAllMovieList()
	c.Data["MovieList"] = movieList
	c.TplName = "index.html"
}
