package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	//var err error
	// 1. 注册数据驱动, mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	// 2. 注册数据库, ORM必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/liteblog?charset=utf8mb4")
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)
	// 根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)

	// 需要在init中注册定义的model
	// 3. 注册模型
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Note))

	// 4. 自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RunSyncdb("default", false, true)

	var (
		db = orm.NewOrm()
	)

	user := User{
		Name:   "admin",
		Email:  "admin@qq.com",
		Pwd:    "admin",
		Avatar: "/static/images/info-img.png",
		Role:   0}

	fmt.Printf("\n---models/core.go--- %v ---\n\n", user)
	//id, err := o.Insert(&user)

	// 判断是否已创建用户，没有则添加
	if created, id, err := db.ReadOrCreate(&user, "Name"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
			fmt.Printf("--- After Insert--- \n---ID: %d, Err: %v ------------\n", id, err)
		}
	}
}

