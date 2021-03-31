package controllers

import (
	"fmt"
	"liteblog/models"
	"liteblog/syserror"
)

type MessageController struct {
	BaseController
}

// @router /message_new/:key [post]
func (this *MessageController) MessageSave() {
	this.MustLogin()
	fmt.Printf("controllers/message.go NextPrepare %v", this.User.Role )

	fmt.Printf("\n\n--- controller/message.go MessageSave() ---  %v  \n\n", this )

	Notekey := this.Ctx.Input.Param(":key")
	//content := this.GetString("content")

	content := this.GetMustString("content","请输入内容！")

	//Notekey := this.GetString("key")
	//content := this.GetString("content")
	fmt.Printf("---controllers/message.go--- NewMessage Notekey:%v, content: %v ", Notekey, content)
	key := this.UUID()

	message1 := &models.Message{
		Key: key,
		NoteKey: Notekey,
		User: this.User.Name,
		UserID: int64(this.User.Id),
		Content: content,
	}

	if err := models.SaveMessage(message1);err!= nil{
		this.Abort500(syserror.New("保存失败",err))
	}
	this.JsonOkH("保存成功！",H{})
	this.TplName = "comment.html"
}