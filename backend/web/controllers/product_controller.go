package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"spike/common"
	"spike/datamodels"
	"spike/services"
	"strconv"
)

type ProductController struct {
	//上下文
	Ctx            iris.Context
	ProductService services.IProductService
}

//显示所有页面
func (p *ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View{Name: "product/view.html", Data: iris.Map{"productArray": productArray}}

}

//添加页面
func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{Name: "product/add.html"}
}

//增加操作
func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	decode := common.NewDecoder(&common.DecoderOptions{TagName: "product"})
	if err := decode.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	if _, err := p.ProductService.InsertProduct(product); err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	p.Ctx.Redirect("all")
}

//修改页面
func (p *ProductController) GetManager() mvc.View {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{Name: "product/manager.html", Data: iris.Map{"product": product}}
}

//修改操作
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
	p.Ctx.Redirect("all")
}

//删除操作
func (p *ProductController) GetDelete() {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	isOk := p.ProductService.DeleteProductByID(id)
	if isOk {
		p.Ctx.Application().Logger().Debug("删除商品成功")
	} else {
		p.Ctx.Application().Logger().Debug("删除商品失败")
	}
	p.Ctx.Redirect("all")
}
