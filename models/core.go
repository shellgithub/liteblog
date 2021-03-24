package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//var (
//	db = orm.NewOrm()
//)

//type User struct {
//    Id int
//	Name string
//	Email string
//	Pwd string
//	Avatar string
//	Role int64 //0代表管理员，1代表正常用户
//}

func init(){
	//var err error

	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/liteblog?charset=utf8mb4")
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)
	// 根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)

	// Create
	//db.Create(&User{
	//	Name:   "admin",
	//	Email:  "admin@qq.com",
	//	Pwd:    "admin",
	//	Avatar: "/static/images/info-img.png",
	//	Role:   0,
	//})

	//如果数据库里没有用户数据，我们新增一条admin 记录
	//var count int64
	db := orm.NewOrm()
	//user := new(User)

	user := User{
		Name:   "admin",
		Email:  "admin@qq.com",
		Pwd:    "admin",
		Avatar: "/static/images/info-img.png",
		Role:   0,}

	//user.Name = "admin"
	//user.Email = "admin@qq.com"
	//user.Pwd = "admin"
	//user.Avatar = "/static/images/info-img.png"
	//user.Role = 1

	fmt.Printf("\n---models/core.go--- %v ---\n\n", user)
	//id, err := o.Insert(&user)


	if created, id, err := db.ReadOrCreate(&user, "Name"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
			fmt.Printf("--- After Insert--- \n---ID: %d, Err: %v ------------\n", id, err)
		}
	}

		//"{
		//Name:   "admin",
		//Email:  "admin@qq.com",
		//Pwd:    "admin",
		//Avatar: "/static/images/info-img.png",
		//Role:   0,}"

	//if created,id, err := o.ReadOrCreate(&user, "Name", "Email", "pwd", "Avatar", "Role") ; err == nil{
	//	if created {
	//		fmt.Println("New Insert an object. Id:", id)
	//	} else {
	//		fmt.Println("Get an object. Id:", id)
	//	}
	//}
	//if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
	//	db.Create(&User{
	//		Name:   "admin",
	//		Email:  "admin@qq.com",
	//		Pwd:    "admin",
	//		Avatar: "/static/images/info-img.png",
	//		Role:   0,
	//	})
	//}


}
