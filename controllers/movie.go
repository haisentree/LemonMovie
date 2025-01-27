package controllers

import (
	"LemonMovie/global"
	"LemonMovie/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
)

type MovieController struct {
	web.Controller
}

func (c *MovieController) GetHome() {
	movieListModel := models.NewMovieModel()
	movieList, maxCount := movieListModel.FindAllMovieList()
	var currentPageID int64 = 1

	maxPage := maxCount/global.PageLimiter + 1
	pageList := []int64{}
	var i int64
	for i = -2; i < 3; i++ {
		if currentPageID+i > 0 && currentPageID+i <= maxPage {
			pageList = append(pageList, currentPageID+i)
		}
	}

	c.Data["CurrentPageID"] = 1
	c.Data["PageList"] = pageList
	c.Data["MovieList"] = movieList
	c.Data["PageName"] = "/page/"
	c.TplName = "index.html"
}

func (c *MovieController) GetHomeByPage() {
	pageID := c.Ctx.Input.Param(":page_id")
	currentIDInt, _ := strconv.ParseInt(pageID, 10, 64)

	movieListModel := models.NewMovieModel()
	movieList, maxCount := movieListModel.FindMovieByPage(currentIDInt)

	maxPage := maxCount/global.PageLimiter + 1
	pageList := []int64{}
	var i int64
	for i = -2; i < 3; i++ {
		if currentIDInt+i > 0 && currentIDInt+i <= maxPage {
			pageList = append(pageList, currentIDInt+i)
		}
	}

	c.Data["CurrentPageID"] = currentIDInt
	c.Data["PageList"] = pageList
	c.Data["PageName"] = "/page/"
	c.Data["MovieList"] = movieList
	c.TplName = "index.html"
}

func (c *MovieController) GetByCategory() {
	typePid := c.Ctx.Input.Param(":id")
	typePidNum, _ := strconv.Atoi(typePid)

	movieListModel := models.NewMovieModel()
	movieList := movieListModel.FindByCategory(typePidNum)

	c.Data["CurrentPageID"] = 1
	c.Data["MovieList"] = movieList
	c.Data["PageName"] = "/category/" + typePid
	c.TplName = "index.html"
}

type PlayVideo struct {
	Name string
	URL  string
}

func (c *MovieController) GetMovieByID() {
	movieID := c.Ctx.Input.Param(":id")
	movieIDNum, _ := strconv.Atoi(movieID)

	movieModel := models.NewMovieModel()
	movieDetail := movieModel.FindByID(movieIDNum)
	// 对播放链接进行处理
	var playDealList []PlayVideo //最终处理完成
	playURL := movieDetail.PlayURL
	// 先以三个$$$进行分割成两部分

	playList := strings.Split(playURL, "$$$")
	temp := playList[1]
	playList2 := strings.Split(temp, "#")
	for _, v := range playList2 {
		tempList := strings.Split(v, "$")
		tempPlay := PlayVideo{
			Name: tempList[0],
			URL:  tempList[1],
		}
		playDealList = append(playDealList, tempPlay)
	}

	c.Data["PlayDealList"] = playDealList
	c.Data["MovieDetail"] = movieDetail
	c.TplName = "movie.html"
}

func (c *MovieController) PlayMovie() {
	movieID := c.Ctx.Input.Param(":movie_id")
	playID := c.Ctx.Input.Param(":play_id")
	movieIDNum, _ := strconv.Atoi(movieID)
	playIDNum, _ := strconv.Atoi(playID)

	movieModel := models.NewMovieModel()
	movieDetail := movieModel.FindByID(movieIDNum)

	// 对播放链接进行处理
	var playDealList []PlayVideo //最终处理完成
	playURL := movieDetail.PlayURL
	// 先以三个$$$进行分割成两部分

	playList := strings.Split(playURL, "$$$")
	temp := playList[1]
	fmt.Println(temp)
	playList2 := strings.Split(temp, "#")
	for _, v := range playList2 {
		tempList := strings.Split(v, "$")
		tempPlay := PlayVideo{
			Name: tempList[0],
			URL:  tempList[1],
		}
		playDealList = append(playDealList, tempPlay)
	}

	c.Data["CurrentPlay"] = playDealList[playIDNum]
	c.Data["PlayDealList"] = playDealList
	c.Data["MovieDetail"] = movieDetail
	c.TplName = "player.html"
}
