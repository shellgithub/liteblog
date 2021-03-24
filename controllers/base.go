package controllers

import (
	"fmt"
	"errors"
	"github.com/beego/beego/v2/server/web"
	"liteblog/models"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	web.Controller
	User models.User
	IsLogin bool
}

func (this *BaseController) Prepare()  {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u,ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
}

func (this *BaseController) Abort500(err error)  {
	this.Data["error"]=err
	this.Abort("500")
}

func (this *BaseController) GetMustString(key ,msg string) string {
	newsKey := this.GetString(key)
	newsKeyLen := len(newsKey)
	fmt.Printf("\n--controllers/base.go newsKey: %v, newsKeyLen:%d -- \n", newsKey, newsKeyLen)
	if len(newsKey) == 0 {
		this.Abort500(errors.New(msg))
	}
	return newsKey
}