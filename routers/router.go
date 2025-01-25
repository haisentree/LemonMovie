package routers

import (
	"LemonMovie/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.CtrlGet("/", (*controllers.MovieController).GetHome)
	web.CtrlGet("/category/:id", (*controllers.MovieController).GetByCategory)
	web.CtrlGet("/movie/:id", (*controllers.MovieController).GetMovieByID)
}
