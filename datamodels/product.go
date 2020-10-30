package datamodels

type Product struct {
	ID int64 `json:"ID" sql:"ID" product:"ID"`
	//商品名称
	ProductName string `json:"ProductName" sql:"productName" product:"ProductName"`
	//商品数量
	ProductNum int64 `json:"ProductNum" sql:"productNum" product:"ProductNum"`
	//商品图片
	ProductImage string `json:"ProductImage" sql:"productImage" product:"Product"`
	//商品地址
	ProductUrl string `json:"ProductUrl" sql:"productUrl" product:"ProductUrl"`
}
