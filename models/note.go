package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)
type Note struct {
	Id     int64
	UserID int64
	Key string
	User string
	Title string
	Summary string
	Content string
	Visit int64
	Praise int64
}

func QueryNoteByKey(key string) (note Note, err error){
	fmt.Printf("\n--models/user.go  user: %v  --\n\n", key)

	var note1 Note
	db := orm.NewOrm()
	qs := db.QueryTable("notes")

	err1 := qs.Filter("key", key).One(&note)
	fmt.Printf("\n--models/note.go  key: %v  --\n\n, err1: %v", key, err1)

	return note1, err1
}

func SaveNote(note *Note) error{
	fmt.Printf("\n--models/user.go 开始执行 SaveUser \n\n " +
		"Key: %v  User: %v  Title: %v , Content: %v --\n\n",
		note.Key, note.User, note.Title, note.Content)
	db := orm.NewOrm()

	fmt.Printf("\n\n note----- %v \n\n", note)
	//err1, _ := db.Insert(users.Id,users.Name,users.Email,users.Pwd,users.Avatar,users.Role)
	id, err := db.Insert(&note)
	if err == nil {
		fmt.Printf("\n\n数据写入成功, id: %d\n\n", id)
	}

	return err
}