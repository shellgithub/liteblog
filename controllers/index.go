package controllers

type IndexController struct {
	BaseController
	//web.Controller
}

//type UserController struct {
//	web.Controller
//}
//
//type MessageController struct {
//	web.Controller
//}

// https://beego.me/docs/mvc/controller/router.md

// @router / [get]
func (this *IndexController) Get() {
	this.TplName = "index.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

//// @router /message [get]
//func (this *IndexController) IndexAbort() {
//	this.Abort("500")
//	this.TplName = "error/500.html"
//}

// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /details [get]
func (this *IndexController) GetDetails() {
	this.TplName = "details.html"
}
