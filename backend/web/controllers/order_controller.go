package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"spike/services"
)

type OrderController struct {
	//上下文
	Ctx          iris.Context
	OrderService services.IOrderService
}

func (o *OrderController) Get() mvc.View {
	orderArray, err := o.OrderService.GetAllOrderWithInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单失败", err)
	}
	return mvc.View{Name: "order/view.html", Data: iris.Map{"orderArray": orderArray}}

}
