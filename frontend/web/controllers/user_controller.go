package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"spike/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (c *UserController) getRegister() mvc.View {
	return mvc.View{Name: "user/register.html"}
}
