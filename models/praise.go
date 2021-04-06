package models

import (
	"fmt"
	orm "github.com/beego/beego/v2/client/orm"
	"liteblog/syserror"
)

type Prasiselog struct {
	Id     int64 //`orm:"index;pk"`
	UserID int64
	Key    string `orm:"size(36)"`
	Table  string `orm:"size(36)"`
	Flag   bool
}



type TempPraise struct {
	Praise int64
}

func UpdatePraise(table,key string, userid int64)(pcnt int64,err error){
	tableType := table

	fmt.Printf("\n\n---models/praise.go UpdatePraise--- table: %v, key: %v, userid: %v\n\n", tableType, key,userid)

	//d := db.Begin()
	//defer func(){
	//	if err!=nil{
	//		d.RollBack()
	//	}else {
	//		d.Commit()
	//	}
	//}
	// 需要判断是否已经点赞过，如果点赞过就发货，点赞过的错
	// 查询点赞流水表，看是否有记录
	var p Prasiselog
	p = Prasiselog{
		Key:     key,
		Table:   tableType,
		UserID:  userid,
		Flag:    false,
	}

	db := orm.NewOrm()
	//qs := db.QueryTable("prasiselog")

	err = db.QueryTable("prasiselog").Filter("table", tableType).Filter("key", key).Filter("UserID", userid).One(&p)

	//created, key, userid, err := db.ReadOrCreate(&p, "key", "UserID") //按字段 key 查询是否含有值为 PrasiseLog.key PrasiseLog.userID 的行，并且会赋值给 p
	if err == nil {
			fmt.Println("\n\n已经点过赞了 num: %v\n\n", err)
			return  0,err
		} else {
			fmt.Printf("\n\n没有点赞，开始写入点赞记录 num: table: %v, key: %v, userid: %v\n\n", tableType, key,userid)
			id, err := db.Insert(&p)
			if err == nil {
				fmt.Printf("Affected id: %v", id)
			}
			fmt.Printf("Affected id: %v", id)
			return 0,err
		}

	if p.Flag {
		return 0,  syserror.HasPraiseError{}
	}


	return pcnt,nil

	//err = d.Model(&PrasiseLog{}).Where("key = ? and table = ? and user_id = ?",key,table,userid ).Take(&p).error
	//if err == gorm.ErrRecordNotFound{
	//	p = PrasiseLog{
	//		UserId: userid,
	//		Key:    key,
	//		Table:  table,
	//		Flag:   false,
	//	}
	//}else if err != nil{
	//	return  0,err
	//}
	//if p.Flag {
	//	return 0,syserror.HasPraiseError{}
	//}


	//p.Flag= true
	// 保存流水
	//if err = d.Save(&p).Error; err!=nil{
	//	return 0,err
	//}

	// 更新文章或留言的点赞数量
	//var (
	//	ppp TempPraise
	//	pcnt int
	//)
	//d.Table(table).Where("key = ?",key).Select("praise").Scan(&ppp).Error
	//if err!=nil{
	//	return 0,err
	//}
	//
	//pcnt = ppp.Praise+1
	//if err := d.Table(table).Where("key = ?", key),Update("praise",pcnt).Error;err!=nil{
	//	return 0,err
	//}
	//
	//return pcnt,nil
}




