package repositories

import (
	"database/sql"
	"spike/common"
	"spike/datamodels"
)

type IProduct interface {
	//连接数据库
	Conn() error
	//插入数据
	Insert(*datamodels.Product) (int64, error)
	//删除数据
	Delete(int64) bool
	//修改数据
	Update(*datamodels.Product) error
	//查询数据
	SelectByKey(int64) (*datamodels.Product, error)
	//查询所有数据
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table, db}
}

//数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMySQLConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

//插入数据
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	if err = p.Conn(); err != nil {
		return 0, err
	}
	insertSql := "insert into product(productName, productNum, productImage, ProductUrl) values(?, ?, ?, ?)"
	stmt, err := p.mysqlConn.Prepare(insertSql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

//删除数据
func (p *ProductManager) Delete(productId int64) bool {
	return true
}

//修改数据
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	return
}

//查询一条数据
func (p *ProductManager) SelectByKey(productId int64) (product *datamodels.Product, err error) {
	return
}

//查询所有数据
func (p *ProductManager) SelectAll() (product []*datamodels.Product, err error) {
	return
}
