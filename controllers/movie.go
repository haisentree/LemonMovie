package controllers

import (
	"LemonMovie/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
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

func (c *MovieController) GetByCategory() {
	typePid := c.Ctx.Input.Param(":id")
	typePidNum, _ := strconv.Atoi(typePid)

	movieListModel := models.NewMovieModel()
	movieList := movieListModel.FindByCategory(typePidNum)
	c.Data["MovieList"] = movieList
	c.TplName = "index.html"
}

func (c *MovieController) GetMovieByID() {
	movieID := c.Ctx.Input.Param(":id")
	movieIDNum, _ := strconv.Atoi(movieID)

	movieModel := models.NewMovieModel()
	movieDetail := movieModel.FindByID(movieIDNum)
	c.Data["MovieDetail"] = movieDetail
	c.TplName = "movie.html"
}
