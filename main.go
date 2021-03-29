package main

import (
	"encoding/gob"
	orm "github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
	"liteblog/models"
	_ "liteblog/routers"
	"strings"

	_ "liteblog/models" // 程序启动时会加载models 目录下的模块
)

func main() {
	initSession()
	initTemplate()
	web.Run()
}

func initSession(){
	gob.Register(models.User{})
	gob.Register(models.Note{})
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Note))

	web.BConfig.WebConfig.Session.SessionOn=true
	web.BConfig.WebConfig.Session.SessionName="liteblog"
	web.BConfig.WebConfig.Session.SessionProvider="file"
	web.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	web.AddFuncMap("equrl", func(x,y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1,y1) == 0
	})

}