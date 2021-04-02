package controllers

import (
	"fmt"
	mod "liteblog/models"
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

// @router /message_new/?:key [post]
func (this *MessageController) MessageSave() {
	fmt.Printf("controllers/message.go MessageSave %v", this.User.Role)

	this.MustLogin()
<<<<<<< HEAD
=======
	fmt.Printf("controllers/message.go NextPrepare %v", this.User.Role)
>>>>>>> 382c245e61573b6f74ef6f55c6906113c507d6da

	fmt.Printf("\n\n--- controller/message.go MessageSave() ---  %v  \n\n", this)

	Notekey := this.Ctx.Input.Param(":key")
	//content := this.GetString("content")

	content := this.GetMustString("content", "请输入内容！")

	//Notekey := this.GetString("key")
	//content := this.GetString("content")
	fmt.Printf("---controllers/message.go--- NewMessage Notekey:%v, content: %v ", Notekey, content)
	key := this.UUID()

<<<<<<< HEAD
	message1 := &models.Message{
=======
	message1 := &mod.Message{
>>>>>>> 382c245e61573b6f74ef6f55c6906113c507d6da
		Key:     key,
		NoteKey: Notekey,
		User:    this.User.Name,
		UserID:  int64(this.User.Id),
		Content: content,
	}

<<<<<<< HEAD
	if err := models.SaveMessage(message1); err != nil {
		this.Abort500(syserror.New("保存失败", err))
	}
	this.JsonOkH("保存成功！", H{"data":message1})
	this.TplName = "comment.html"
}

=======
	if err := mod.SaveMessage(message1); err != nil {
		this.Abort500(syserror.New("保存失败", err))
	}
	this.JsonOkH("保存成功！", H{})
	this.TplName = "comment.html"
}
>>>>>>> 382c245e61573b6f74ef6f55c6906113c507d6da
