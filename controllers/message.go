package controllers

import (
	"fmt"
	"liteblog/models"
	"liteblog/syserror"
)

type MessageController struct {
	BaseController
}

// @router /message/count [get]
func (this *MessageController) Count () {
	count, err := models.QueryMessageCountByNoteKey("")
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功！", H{"count":count})
}

// @router /message/query [get]
func (this *MessageController) Query () {
	pageno, err:= this.GetInt("pageno")
	if err != nil || pageno <1 {
		pageno = 1
	}
	pagesize, err := this.GetInt("pagesize")
	if err != nil{
		pagesize = 5
	}

	ms, err := models.QueryPageMessageByNoteKey("",pageno,pagesize)
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功！", H{"data":ms})
}

// @router /message_new/ [post]
func (this *MessageController) MessageSave() {
	fmt.Printf("controllers/message.go MessageSave %v", this.User.Role)

	this.MustLogin()

	fmt.Printf("\n\n--- controller/message.go MessageSave() ---  %v  \n\n", this)

	Notekey := this.Ctx.Input.Param(":key")
	//content := this.GetString("content")

	content := this.GetMustString("content", "请输入内容！")

	//Notekey := this.GetString("key")
	//content := this.GetString("content")
	fmt.Printf("---controllers/message.go--- NewMessage Notekey:%v, content: %v ", Notekey, content)
	key := this.UUID()

	message1 := &models.Message{
		Key:     key,
		NoteKey: Notekey,
		User:    this.User.Name,
		UserID:  int64(this.User.Id),
		Content: content,
	}

	if err := models.SaveMessage(message1); err != nil {
		this.Abort500(syserror.New("保存失败", err))
	}
	this.JsonOkH("保存成功！", H{"data":message1})
	this.TplName = "comment.html"
}

