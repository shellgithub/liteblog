package controllers

import (
	"fmt"
	"liteblog/models"
	"liteblog/syserror"
)

type IndexController struct {
	BaseController
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
	var (
		limit int
		count int64
	)
	limit = 4
	//page
	page, err := this.GetInt("page",1)
	if err != nil || page <=0 {
		page=1
	}
	fmt.Printf("\n---index.go ---page: %v, limit: %v\n",page, limit)

	title := this.GetString("title")
	// 得到的页面的数据
	notes, err := models.QueryNotesByPage(title, page, limit)
	if err!= nil{
		this.Abort500(err)
	}
	this.Data["notes"] = notes

	// 得到的页面数量
	count, err = models.QueryNotesCount(title)
	if err!= nil{
		this.Abort500(err)
	}
	fmt.Printf("\n\n---controlleres/index.go 文章总数, %v---\n\n", count)

	totpage := count/int64(limit)
	if count % int64(limit) != 0 {
		totpage = totpage + 1
	}
	this.Data["totpage"] = totpage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) GetReg() {
	this.TplName = "reg.html"
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

// @router /comment [get]
func (this *IndexController) GetComment() {
	key := this.GetString("key")
	fmt.Printf("---controllers/index.go---GetDetails key: %v\n",key)

	note,err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["note"] = note

	// 评论的页面的数据
	messages ,err := models.QueryMessageByNoteKey(key)
	if err != nil {
		this.Abort500(syserror.New("评论不存在",err))
	}
	this.Data["messages"] = messages

	this.TplName = "comment.html"
}

// @router /details [get]
func (this *IndexController) GetDetails() {
	key := this.GetString("key")
	fmt.Printf("---controllers/index.go---GetDetails key: %v\n",key)

	note,err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["note"] = note

	ms, err := models.QueryMessageByNoteKey(key)
	if err != nil {
		this.Abort500(syserror.New("评论不存在",err))
	}
	this.Data["messages"] = ms
	this.TplName = "details.html"
}

