package repositories

import (
	"database/sql"
	"spike/common"
	"spike/datamodels"
	"strconv"
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
	insertSql := "insert into " + p.table + "(productName, productNum, productImage, productUrl) values(?, ?, ?, ?)"
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
func (p *ProductManager) Delete(productID int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}
	deleteSql := "delete from " + p.table + " where ID=?"
	stmt, err := p.mysqlConn.Prepare(deleteSql)
	if err != nil {
		return false
	}
	if _, err = stmt.Exec(productID); err != nil {
		return false
	}
	return true
}

//修改数据
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	if err := p.Conn(); err != nil {
		return err
	}
	updateSql := "update " + p.table + " set ProductName=?, ProductNum=?, productImage=?, productUrl=? where ID=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(updateSql)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl); err != nil {
		return err
	}
	return
}

//查询一条数据
func (p *ProductManager) SelectByKey(productID int64) (product *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}
	selectOneSql := "select * from " + p.table + " where ID=" + strconv.FormatInt(productID, 10)
	row, err := p.mysqlConn.Query(selectOneSql)
	if err != nil {
		return &datamodels.Product{}, err
	}
	result := common.GetResultRow(row)
	product = &datamodels.Product{}
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	//数据映射到结构体
	common.DataToStructByTagSql(result, product)
	return
}

//查询所有数据
func (p *ProductManager) SelectAll() (productArray []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return nil, err
	}
	selectAllSql := "select * from " + p.table
	rows, err := p.mysqlConn.Query(selectAllSql)
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	//数据映射到结构体
	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return
}
