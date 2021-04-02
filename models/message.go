package models

import (
	"fmt"
	orm "github.com/beego/beego/v2/client/orm"
	"github.com/goinggo/mapstructure"
	"time"
)

type Message struct {
	Id      int64
	Key     string    `orm:"size(36)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Content string    `orm:"type(text)"`
	UserID  int64
	User    string `orm:"size(50)"`
	NoteKey string `orm:"size(36)"`
	Praise  int64  `default:"0"`
}

func SaveMessage(message *Message) error {
	fmt.Printf("\n--models/message.go 开始执行 SaveNote \n\n "+
		"Key: %v, Userid: %v,  User: %v  Title: %v , Content: %v --\n\n",
		message.Key, message.UserID, message.User, message.NoteKey, message.Content)

	message1 := Message{
		Key:     message.Key,
		User:    message.User,
		UserID:  message.UserID,
		NoteKey: message.NoteKey,
		Content: message.Content,
	}
	Key := message1.Key

	db := orm.NewOrm()

	fmt.Printf("\n\n ----接收到的message内容： %v ---\n\n", message)
	fmt.Printf("\n\n ----接收到的message1内容： %v ---\n\n", message1)
	fmt.Printf("\n\n ----接收到的message1内容 Key: %v, Title: %v, Content: %v, 开始更新message 内容\n\n",
		message1.Key, message1.UserID, message1.Content)
	fmt.Printf("\n\n ----开始写入到库---\n\n")
	created, key, err := db.ReadOrCreate(&message1, "key") //按字段 key 查询是否含有值为 note1.key 的行，并且会赋值给 note1
	if err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", key)
		} else {
			fmt.Printf("Get an object. Id: %v, 开始更新message 内容\n\n", Key)
			fmt.Printf("Get an object. Key: %v, User: %v, Userid: %v, NoteKey: %v, Content: %v, 开始更新message 内容\n\n",
				message.Key, message.User, message.UserID, message.NoteKey, message.Content)
			num, err := db.QueryTable("message").Filter("key", Key).Update(orm.Params{
				"Key":     message.Key,
				"User":    message.User,
				"UserID":  message.UserID,
				"Content": message.Content,
				"NoteKey": message.NoteKey,
				"Updated": time.Now(),
			})
			fmt.Printf("Affected Num: %v, %v", num, err)
		}
	}
	return err
}

func QueryMessageByNoteKey(notekey string) (message []orm.Params, err error) {

	fmt.Printf("\n--models/message.go  notekey: %v  --\n\n", notekey)
	db := orm.NewOrm()

	var maps []orm.Params
	_, err = db.QueryTable("message").Filter("NoteKey", notekey).OrderBy("-updated").Values(&maps)

	if err := mapstructure.Decode(maps, &message); err != nil {
		fmt.Println(err)
	}

	return maps, err
}

//func QueryNotesByPage11(title string, page , limit int) (note []orm.Params,err error){
//	firstLimit := (page-1)*limit
//	//limit := page*limit
//	fmt.Printf("---models/note.go--- QueryNotesByPage  firstLimit:%v , limit: %v ",firstLimit, limit)
//
//	db := orm.NewOrm()
//	// 获取 QuerySeter 对象，user 为表名
//	var maps []orm.Params
//	_, err = db.QueryTable("note").Filter("title__icontains", title).OrderBy("-updated").Limit(limit, firstLimit).Values(&maps)
//
//	if err := mapstructure.Decode(maps, &note); err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Printf("\n---maps ---%v\n",maps)
//	fmt.Printf("\n---note ---%v\n",note)
//	return maps, err
//}
