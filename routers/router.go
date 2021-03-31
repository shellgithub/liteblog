package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"liteblog/controllers"
)

// https://beego.me/docs/mvc/controller/router.md
//注解路由
//从2.0开始，我们使用配置CommentRouterPath来配置注解路由的扫描路径。在dev环境下，我们将自动扫描该配置指向的目录及其子目录，生成路由文件。
//生成之后，用户需要显示 Include 相应地 controller。注意， controller 的 method 方法上面须有 router 注释（// @router）

func init() {

	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.NoteController{})
	beego.Include(&controllers.MessageController{})

    // 使用 note/new 地址 静态资源加载不了
	//beego.AddNamespace(
	//	beego.NewNamespace(
	//		"/note",
	//		beego.NSInclude(&controllers.NoteController{}),
	//	),
	//)

}
