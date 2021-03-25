package controllers

import (
	"errors"
	"fmt"
	"liteblog/models"
	"liteblog/syserror"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	//email
	email := this.GetMustString("email", "邮箱不能为空！")
	fmt.Printf("\n--controllers/user.go-- email: %v\n ", email)

	//password
	pwd := this.GetMustString("password", "密码不能为空！")
	fmt.Printf("\n--controllers/user.go-- pwd: %v\n ", pwd)

	user, err := models.QueryUserByEmailAndPwd(email, pwd)
	fmt.Printf("\n--controllers/user.go-- pwd: %v\n ", user)

	//fmt.Print("user", user)

	if err != nil {
		fmt.Print("err", err)
		this.Abort500(syserror.New("登录失败！", err))
	}
	this.SetSession(SESSION_USER_KEY, user)
	this.JsonOk("登录成功", "/")
}

// @router /reg [post]
func (this *UserController) Reg() {
	// 昵称、邮箱、密码、确认密码 都不能为空
	//name
	name := this.GetMustString("name", "昵称不能为空！")
	email := this.GetMustString("email", "邮箱不能为空！")
	password := this.GetMustString("password", "密码不能为空！")
	password2 := this.GetMustString("password1", "确认密码不能为空！")

	fmt.Printf("\n--controllers/user.go--UserController Reg \n"+
		"name: %v,email: %v,password: %v,password2: %v\n ", name, email, password, password2)

	fmt.Print("\n开始判断两次输入的密码是否一致\n")

	if strings.Compare(password, password2) != 0 {
		//msg := "两次输入的密码不一致"
		this.Abort500(errors.New("不一致"))
		//this.Abort500(errors.New(msg))
		//fmt.Println("两次输入的密码不一致")
	} else {
		fmt.Printf("\n两次输入的密码一致\n")
	}
	fmt.Print("\n结束判断两次输入的密码是否一致\n")
	fmt.Print("\n开始判断用户昵称是否已存在\n")

	if u, err := models.QueryUserByName(name); err == nil && u.Id > 0 {
		this.Abort500(errors.New("用户昵称已存在"))
	} else {
		fmt.Print("\n用户昵称不存在\n")
	}

	if u, err := models.QueryUserByEmail(email); err == nil && u.Id > 0 {
		this.Abort500(errors.New("用户邮箱已存在"))
	} else {
		fmt.Print("\n用户邮箱不存在\n")
	}

	fmt.Printf("\n--controllers/user.go--name: %v,email: %v,password: %v,password2: %v\n ", name, email, password, password2)
	fmt.Printf("\n\n\n--controllers/user.go--调用models.SaveUser 开始保存数据到库\n\n\n")

	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Pwd:    password,
		Avatar: "/static/images/info-img.png",
		Role:   0,
	}); err != nil {
		this.Abort500(syserror.New("用户保存失败", err))
	}
	this.JsonOk("注册成功", "/user")
}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}
