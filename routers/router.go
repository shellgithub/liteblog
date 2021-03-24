package routers

import (
	"github.com/beego/beego/v2/server/web"
	"liteblog/controllers"
)

// https://beego.me/docs/mvc/controller/router.md

func init() {

	web.ErrorController(&controllers.ErrorController{})
	web.Include(&controllers.IndexController{})
	web.Include(&controllers.UserController{})

}
