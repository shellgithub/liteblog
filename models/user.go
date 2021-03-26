package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	//"gorm.io/gorm"
)

// 需要在主程序中添加import _ "liteblog/models"  // 即程序启动时会加载models 目录下的模块

//type User struct {
//	//gorm.Model
//	Name string `gorm:"unique_index"`
//	Email string `gorm:"unique_index"`
//	Pwd string
//	Avatar string
//	Role int64 `gorm:"default:1"` //0代表管理员，1代表正常用户
//}

type User struct {
	Id     int
	Name   string
	Email  string
	Pwd    string
	Avatar string
	Role   int64 //0代表管理员，1代表正常用户
}

func QueryUserByEmailAndPwd(email string, pwd string) (user User, err error) {
	//defer
	fmt.Printf("\n--models/user.go email: %v, pwd: %v --\n\n", email, pwd)

	var user1 User
	db := orm.NewOrm()
	qs := db.QueryTable("user")
	//user := User{Name: email, Pwd: pwd}
	err1 := qs.Filter("email", email).Filter("pwd", pwd).One(&user)
	fmt.Printf("\n--models/user.go  user: %v  --\n\n, err1: %v", user1, err1)

	return user, err1

}

func QueryUserByName(name string) (user User, err error) {
	fmt.Printf("\n--models/user.go  user: %v  --\n\n", user)

	//var user1 User
	db := orm.NewOrm()
	qs := db.QueryTable("user")

	err1 := qs.Filter("name", name).One(&user)
	fmt.Printf("\n--models/user.go  user: %v  --\n\n, err1: %v", user, err1)

	return user, err1
}

func QueryUserByEmail(email string) (user User, err error) {
	var user1 User
	db := orm.NewOrm()
	qs := db.QueryTable("user")

	err1 := qs.Filter("email", email).One(&user)
	fmt.Printf("\n--models/user.go  user: %v  --\n\n, err1: %v", user1, err1)

	return user, err1
}

func SaveUser(user *User) error {
	fmt.Printf("\n--models/user.go 开始执行 SaveUser \n\n"+
		"Email: %v  Name: %v , Pwd: %v --\n\n", user.Email, user.Name, user.Pwd)
	db := orm.NewOrm()

	fmt.Printf("\n\nusers----- %v \n\n", user)
	//err1, _ := db.Insert(users.Id,users.Name,users.Email,users.Pwd,users.Avatar,users.Role)
	id, err := db.Insert(user)
	if err == nil {
		fmt.Printf("\n\n数据写入成功, id: %d\n\n", id)
	}

	return err
}
