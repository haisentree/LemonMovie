package routers

import (
	"LemonMovie/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.CtrlGet("/", (*controllers.MovieController).GetHome)
	web.CtrlGet("/page/:page_id", (*controllers.MovieController).GetHomeByPage)
	web.CtrlGet("/category/:id", (*controllers.MovieController).GetByCategory)
	web.CtrlGet("/movie/:id", (*controllers.MovieController).GetMovieByID)
	web.CtrlGet("/player/:movie_id/:play_id", (*controllers.MovieController).PlayMovie)
}
