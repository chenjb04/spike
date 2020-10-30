package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"spike/backend/web/controllers"
	"spike/common"
	"spike/repositories"
	"spike/services"
)

func main() {
	//创建iris
	app := iris.New()
	//设置错误等级
	app.Logger().SetLevel("debug")
	//注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置静态资源目录
	app.HandleDir("/assets", "./backend/web/assets")
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//注册控制器
	productRepository := repositories.NewProductManager("product", db)
	ProductService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, ProductService)
	product.Handle(new(controllers.ProductController))
	//启动服务
	app.Run(iris.Addr("localhost:8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)

}
