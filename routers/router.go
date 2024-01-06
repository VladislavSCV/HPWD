package routers

import (
	"hww/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/main_page", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthController{})
	beego.Router("/register", &controllers.RegController{})
	beego.Router("/news_1", &controllers.NewsController{}, "get:News1")
	beego.Router("/news_2", &controllers.NewsController{}, "get:News2")
	beego.Router("/news_3", &controllers.NewsController{}, "get:News3")
	beego.Router("/job_openings", &controllers.CatalogController{})
}