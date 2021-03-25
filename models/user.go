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
    Id int
	Name string
	Email string
	Pwd string
	Avatar string
	Role int64 //0代表管理员，1代表正常用户
}

func QueryUserByEmailAndPwd(email string, pwd string) (user User,err error) {
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
