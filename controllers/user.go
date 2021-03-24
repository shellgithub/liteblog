package controllers

import (
	"fmt"
	"liteblog/models"
	"liteblog/syserror"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	//email
	email := this.GetMustString("email","邮箱不能为空！")
    fmt.Printf("\n--controllers/user.go-- email: %v\n ", email)

	//password
	pwd := this.GetMustString("password","密码不能为空！")
	fmt.Printf("\n--controllers/user.go-- pwd: %v\n ", pwd)


	user, err := models.QueryUserByEmailAndPwd(email, pwd)
	fmt.Printf("\n--controllers/user.go-- pwd: %v\n ", user)

	//fmt.Print("user", user)

	if err != nil {
		fmt.Print("err", err)
		this.Abort500(syserror.New("登录失败！", err))
	}
	this.SetSession(SESSION_USER_KEY, user)
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"action": "/",
	}
	//fmt.Print(SESSION_USER_KEY, user)

	this.ServeJSON()

}