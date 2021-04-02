package controllers

import (
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	//"go/build/constraint"
	"liteblog/models"
	"liteblog/syserror"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool
}

type NestPreparer interface {
	NextPrepare()
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["islogin"] = this.IsLogin
	if a, ok := this.AppController.(NestPreparer); ok {
		a.NextPrepare()
	}
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) GetMustString(key, msg string) string {
	newsKey := this.GetString(key)
	newsKeyLen := len(newsKey)
	fmt.Printf("\n--controllers/base.go-GetMustString-- newsKey: %v, newsKeyLen:%d -- \n", newsKey, newsKeyLen)
	if len(newsKey) == 0 {
		this.Abort500(errors.New(msg))
	}
	return newsKey
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

type H map[string]interface{}

func (this *BaseController) JsonOk(msg, action string) {
	this.Data["json"] = H{
		"code":   0,
		"msg":    msg,
		"action": "/",
	}
	this.ServeJSON()
}

func (this *BaseController) JsonOkH(msg string, data H) {
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()

}

func (this *BaseController) UUID() string {
	u := uuid.NewV4()
	return u.String()
}
