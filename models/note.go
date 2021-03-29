package models

import (
	"bytes"
	"fmt"
	orm "github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/PuerkitoBio/goquery"
	"github.com/goinggo/mapstructure"
	"time"
)
type Note struct {
	Id     int64 `orm:"index"`
	UserID int64
	Key string `orm:"size(36)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	User string `orm:"size(50)"`
	Title string `orm:"size(500)"`
	Summary string `orm:"size(600)"`
	Content string `orm:"type(text)""`
	Visit int64 `default:"0"`
	Praise int64 `default:"0"`
}

//QueryNoteByKey
func QueryNoteByKey(key string) (note Note, err error){
	fmt.Printf("\n--models/user.go  user: %v  --\n\n", key)

	var note1 Note
	db := orm.NewOrm()
	qs := db.QueryTable("note")

	err1 := qs.Filter("key", key).One(&note)
	if err1 != nil {
		fmt.Printf("\n--models/note.go 没查到 key: %v  --\n\n, err1: %v", key, err1)
	}else {
		fmt.Printf("\n--models/note.go 查到有 key: %v  --\n\n, err1: %v", key, err1)
	}
	return note1, err1
}

func QueryNotesByPage(page , limit int) (note []orm.Params,err error){
	firstLimit := (page-1)*limit
	//limit := page*limit
	fmt.Printf("---models/note.go--- QueryNotesByPage  firstLimit:%v , limit: %v ",firstLimit, limit)

	db := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	var maps []orm.Params
	_, err = db.QueryTable("note").OrderBy("-id").Limit(limit, firstLimit).Values(&maps)
	//if err == nil {
	//	fmt.Printf("Result Nums: %d\n", num)
	//	for _, m := range maps{
	//		//fmt.Printf("\n---models/note.go  notes的类型 %T %v ---\n", notes, notes)
	//		fmt.Printf("---models/note.go  QueryNotesByPage id: %v Title: %v\n",m["Id"],m["Title"])
	//	}
	//}
	if err := mapstructure.Decode(maps, &note); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n---maps ---%v\n",maps)
	fmt.Printf("\n---note ---%v\n",note)
	return maps, err
}

func QueryNotesCount() (count int64, err error) {
	db := orm.NewOrm()
	var cnt int64
	cnt, err = db.QueryTable("note").Count()
	return cnt, err

}

//SaveNote
func SaveNote(note *Note) error{
	fmt.Printf("\n--models/note.go 开始执行 SaveNote \n\n " +
		"Key: %v  User: %v  Title: %v , Content: %v --\n\n",
		note.Key, note.User, note.Title, note.Content)

	note1 := Note{
		Key: note.Key,
		User: note.User,
		Title: note.Title,
		Summary: note.Summary,
		Content: note.Content,
	}
	Key := note1.Key
	db := orm.NewOrm()

	fmt.Printf("\n\n ----接收到的note内容： %v ---\n\n", note)
	fmt.Printf("\n\n ----接收到的note内容 Key: %v, Title: %v, Content: %v, 开始更新note 内容\n\n", Key, note1.Title, note1.Content)

	//err1, _ := db.Insert(users.Id,users.Name,users.Email,users.Pwd,users.Avatar,users.Role)
	fmt.Printf("\n\n ----开始写入到库---\n\n")


	created, key, err := db.ReadOrCreate(&note1, "key")  //按字段 key 查询是否含有值为 note1.key 的行，并且会赋值给 note1
	if err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", key)
		} else {
			fmt.Printf("Get an object. Id: %v, 开始更新note 内容\n\n", Key)
			fmt.Printf("Get an object. Title: %v, Summary: %v, Content: %v, 开始更新note 内容\n\n", note.Title,note.Summary,note.Content)
			num, err := db.QueryTable("note").Filter("key", Key).Update(orm.Params{
				"Title": note.Title,
				"Content": note.Content,
				"Summary": note.Summary,
			})
			fmt.Printf("Affected Num: %v, %v", num, err)
		}
	}
	return err

	//id, err := db.Insert(note)
	//if err == nil {
	//	fmt.Printf("\n\n数据写入成功, id: %d\n\n", id)
	//}else{
	//	fmt.Printf("\n\n数据写入失败, err: %v\n\n", err)
	//
	//}
	//return err
}

func GetSummary(html string) (string, error) {
	var bufbytes bytes.Buffer

	if _,err := bufbytes.Write([]byte(html));err!= nil{
		return "",err
	}
	doc, err:= goquery.NewDocumentFromReader(&bufbytes)
	if err!=nil{
		return  "",err
	}
	htmlstr := doc.Find("body").Text()
	if len(htmlstr)>600{
		htmlstr = htmlstr[:600]
	}
	return htmlstr,nil
}
