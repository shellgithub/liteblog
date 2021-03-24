package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"

	//"gorm.io/gorm"
)

// 需要在主程序中添加import _ "liteblog/models"  // 即程序启动时会加载models 目录下的模块
//type UserStruct struct {
//	//gorm.Model
//	Name string `gorm:"unique_index"`
//	Email string `gorm:"unique_index"`
//	Pwd string
//	Avatar string
//	Role int64 `gorm:"default:1"`
//}

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

	//user1 := db.Where("email = ? and pwd = ?", email, pwd).Take(&users)
	//db := orm.NewOrm()
	//r := db.Raw("select * from users where email = ? and pwd = ?")
	//res, err1:= r.SetArgs(email, pwd).Exec()
	//if err1 != nil {
	//	//res.RowsAffected()
	//	fmt.Printf("查询成功！")
	//}else{
	//	fmt.Printf("查询失败！")
	//}
	//
	//fmt.Printf("models/user.go,res  %v\n", res)
	//return user, err1

	// 获取第一条记录（主键升序）
	//var User1 User
	//result := db.First(&user)
	//UserCount := result.RowsAffected // 返回找到的记录数
	//fmt.Printf("\n --models/user.go  UserCount: %v \n\n", UserCount)
	//
	//err = result.Error        // returns error
	//fmt.Printf("\n --models/user.go  UserCount: %v err: %v \n\n", UserCount, err)
	//
	//var result1 = db.Where("email = ? and pwd = ?", email, pwd).Take(&user)
	//var UserCount1 = result1.RowsAffected // 返回找到的记录数
	//var err1 = result1.Error              // returns error
	//fmt.Printf("\n --models/user.go  UserCount: %v err: %v \n\n", UserCount1, err1)
	//
	//err = db.Where("email = ? and pwd = ?", email, pwd).Take(&user).Error
	//if err != nil {
	//	fmt.Print( UserCount, user, err)
	//	return user, err
	//}else {
	//	fmt.Print( UserCount, user, err)
	//	return user, err
	//}

	//return user, db.Where("email = ? and pwd = ?", email, pwd).Find(&user).Error


}
