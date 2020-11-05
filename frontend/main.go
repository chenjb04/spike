package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"log"
	"spike/common"
	"spike/frontend/web/controllers"
	"spike/repositories"
	"spike/services"
	"time"
)

func main() {
	//创建iris
	app := iris.New()
	//设置错误等级
	app.Logger().SetLevel("debug")
	//注册模板
	template := iris.HTML("./frontend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置静态资源目录
	app.HandleDir("/public", "./frontend/web/public")
	//出现异常跳转
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")

	})
	db, err := common.NewMySQLConn()
	if err != nil {
		log.Fatal(err)
	}
	session := sessions.New(sessions.Config{Cookie: "helloworld", Expires: 60 * time.Minute})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//注册控制器
	userRepository := repositories.NewUserManagerRepository("user", db)
	userService := services.NewUserService(userRepository)
	user := mvc.New(app.Party("/user"))
	user.Register(userService, ctx, session.Start)
	user.Handle(new(controllers.UserController))
	//启动服务
	app.Run(iris.Addr("0.0.0.0:8082"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)

}
