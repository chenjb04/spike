package main

import "github.com/kataras/iris"

func main() {
	//创建iris
	app := iris.New()
	//设置错误等级
	app.Logger().SetLevel("debug")
	//注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置静态资源目录
	app.HandleDir("/assets", iris.Dir("./backend/web/assets"))
	//出现异常跳转
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")

	})
	//注册控制器
	//启动服务
	app.Run(iris.Addr("localhost:8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)

}
