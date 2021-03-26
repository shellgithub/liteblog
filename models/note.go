package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)
type Note struct {
	Key string
	UserID int64
	User string
	Title string
	Summary string
	Content string
	Visit int
	Praise int
}

func QueryNoteByKey(key string) (note Note, err error){
	fmt.Printf("\n--models/user.go  user: %v  --\n\n", key)

	//var user1 User
	db := orm.NewOrm()
	qs := db.QueryTable("note")

	err1 := qs.Filter("key", key).One(&note)
	fmt.Printf("\n--models/note.go  user: %v  --\n\n, err1: %v", key, err1)

	return note, err1
}

func SaveNote(note *Note) error{
	fmt.Printf("\n--models/user.go 开始执行 SaveUser \n\n"+
		"Key: %V  User: %v  Title: %v , Content: %v --\n\n", note.Key, note.User, note.Title, note.Content)
	db := orm.NewOrm()

	fmt.Printf("\n\n note----- %v \n\n", note)
	//err1, _ := db.Insert(users.Id,users.Name,users.Email,users.Pwd,users.Avatar,users.Role)
	id, err := db.Insert(note)
	if err == nil {
		fmt.Printf("\n\n数据写入成功, id: %d\n\n", id)
	}

	return err
}