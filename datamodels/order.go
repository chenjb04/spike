package datamodels

type Order struct {
	//订单编号
	ID int64 `json:"ID" sql:"ID"`
	//用户id
	UserID int64 `json:"UserID" sql:"userID"`
	//商品ID
	ProductId int64 `json:"ProductId" sql:"productId"`
	//订单状态
	OrderStatus int64 `json:"OrderStatus" sql:"orderStatus"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)
