package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"spike/common"
	"spike/datamodels"
	"spike/services"
)

type ProductController struct {
	//上下文
	Ctx            iris.Context
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View{Name: "product/view.html", Data: iris.Map{"productArray": productArray}}

}

func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	decode := common.NewDecoder(&common.DecoderOptions{TagName: "product"})
	if err := decode.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	if err := p.ProductService.UpdateProduct(product); err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	p.Ctx.Redirect("product/all")
}
