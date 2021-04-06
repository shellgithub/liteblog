package controllers

import (
	"errors"
	"liteblog/models"
	"liteblog/syserror"

)

type PraiseController struct {
	BaseController
}

func (this *PraiseController) NextPrepare(){
	this.MustLogin()
}

// @router /praise/:type/:key [post]
func (this *PraiseController) Praise(){
	ttype := this.Ctx.Input.Param(":type")
	key := this.Ctx.Input.Param(":key")
	userid := int64(this.User.Id)

	table := "note"
	switch ttype {
	case "message":
		table = "message"
	case "note":
		table = "note"
	default:
		this.Abort500(errors.New("未知类型"))
	}


	pcnt, err := models.UpdatePraise(table ,key , userid )
	if err != nil{
		if e2,ok := err.(syserror.HasPraiseError); ok{
			this.Abort500(e2)
		}
		this.Abort500(syserror.New("点赞失败", err))

	}

	this.JsonOkH("点赞成功",H{"praise":pcnt})

}