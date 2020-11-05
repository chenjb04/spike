package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"spike/datamodels"
	"spike/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

//注册页面
func (c *UserController) GetRegister() mvc.View {
	return mvc.View{Name: "user/register.html"}
}

//注册操作
func (c *UserController) PostRegister() {
	//user := &datamodels.User{}
	//c.Ctx.Request().ParseForm()
	//decode := common.NewDecoder(&common.DecoderOptions{TagName: "user"})
	//if err := decode.Decode(c.Ctx.Request().Form, user); err != nil {
	//	c.Ctx.Application().Logger().Error(err)
	//}

	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		pwd      = c.Ctx.FormValue("password")
	)
	user := &datamodels.User{
		NickName:     nickName,
		UserName:     userName,
		HashPassword: pwd,
	}
	if _, err := c.Service.AddUser(user); err != nil {
		c.Ctx.Application().Logger().Error(err)
	}
	c.Ctx.Redirect("login")
	return
}
