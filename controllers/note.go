package controllers

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	. "liteblog/models"
	"liteblog/syserror"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) NextPrepare(){
	this.MustLogin()
	fmt.Printf("controllers/note.go NextPrepare %v", this.User.Role )
	if this.User.Role != 0 {
		fmt.Printf("controllers/note.go NextPrepare %v", this.User.Role )
		this.Abort500(errors.New("权限不足"))
	}
}

// /note
// @router /note_new [get]
func (this *NoteController) NoteNew() {
	this.Data["key"] = this.UUID()
	this.TplName = "note_new.html"
}

// @router /note_save/:key [post]
func (this *NoteController) NoteSave() {
	fmt.Printf("\n\n--- controller/note.go NoteSave() ---  %v  \n\n", "this" )

	key := this.Ctx.Input.Param(":key")
	fmt.Printf("\n\n--- controller/note.go NoteSave() 获取从页面传过来 key 的值---  %v  \n\n", key )
	// title content
	title := this.GetMustString("title", "请输入标题")
	content := this.GetMustString("content", "请输入文章内容")

	fmt.Printf("\n\n--- controller/note.go NoteSave() --- \n\n key: %v  \n " +
		"title: %v\n content: %v\n", key, title, content )

	var n Note
	note, err := QueryNoteByKey(key)
	if err!=nil{
		if err == orm.ErrNoRows{
			n = Note{
				Key:     key,
				Title:   title,
				Content: content,
				UserID: int64(this.User.Id),
				User: this.User.Name,
			}
		}else {
			this.Abort500(syserror.New("保存失败",err))
		}
	}else {
		note.Title = title
		note.Content = content
		n = Note{
			Key:     key,
			Title:   title,
			Content: content,
			UserID: int64(this.User.Id),
			User: this.User.Name,
		}
	}

	// 截取摘要
	n.Summary, _ = GetSummary(content)


	if err := SaveNote(&n); err != nil {
		this.Abort500(syserror.New("保存失败",err))
	}
	this.JsonOkH("保存成功",H{})
	this.TplName = "note_new.html"
}

