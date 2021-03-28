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
	key := this.Ctx.Input.Param(":key")
	fmt.Printf("\n\n--- controller/note.go NoteSave() ---  %v  \n\n", key )
	// title content
	title := this.GetMustString("title", "请输入标题")
	content := this.GetMustString("content", "请输入文章内容")

	fmt.Printf("\n\n--- controller/note.go NoteSave() --- \n\n %v  \n title: %v\n content: %v\n", key,title,content )

	note, err := QueryNoteByKey(key)
	var n Note
	if err!=nil{
		if err == orm.ErrNoRows{
			n = Note{
				Key:     key,
				Title:   title,
				Content: content,
				UserID: int64(this.User.Id),
				//User: this.User.Name,
			}
			//models.SaveNote(&note)
		}else {
			this.Abort500(syserror.New("保存失败",err))
		}
	}else {
		note.Title = title
		note.Content = content
		n = note
	}
	//models.SaveNote(&note)
	if err := SaveNote(&n); err != nil {
		this.Abort500(syserror.New("保存失败",err))
	}
	this.TplName = "note_new.html"
}
