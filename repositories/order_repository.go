package repositories

import (
	"database/sql"
	"spike/common"
	"spike/datamodels"
	"strconv"
)

type IOrder interface {
	//连接数据库
	Conn() error
	//插入数据
	Insert(*datamodels.Order) (int64, error)
	//删除数据
	Delete(int64) bool
	//修改数据
	Update(*datamodels.Order) error
	//查询数据
	SelectByKey(int64) (*datamodels.Order, error)
	//查询所有数据
	SelectAll() ([]*datamodels.Order, error)
	//查询订单消息
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewOrderManager(table string, db *sql.DB) IOrder {
	return &OrderManager{table, db}
}

//数据库连接
func (o *OrderManager) Conn() (err error) {
	if o.mysqlConn == nil {
		mysql, err := common.NewMySQLConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return
}

//插入数据
func (o *OrderManager) Insert(order *datamodels.Order) (orderId int64, err error) {
	if err = o.Conn(); err != nil {
		return 0, err
	}
	insertSql := "insert into " + o.table + "(userID, productID, OrderStatus) values(?, ?, ?)"
	stmt, err := o.mysqlConn.Prepare(insertSql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(order.UserID, order.ProductId, order.OrderStatus)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

//删除数据
func (o *OrderManager) Delete(orderID int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}
	deleteSql := "delete from " + o.table + " where ID=?"
	stmt, err := o.mysqlConn.Prepare(deleteSql)
	if err != nil {
		return false
	}
	if _, err = stmt.Exec(orderID); err != nil {
		return false
	}
	return true
}

//修改数据
func (o *OrderManager) Update(order *datamodels.Order) (err error) {
	if err := o.Conn(); err != nil {
		return err
	}
	updateSql := "update " + o.table + " set userID=?, productID=?,OrderStatus=?where ID=" + strconv.FormatInt(order.ID, 10)
	stmt, err := o.mysqlConn.Prepare(updateSql)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(order.UserID, order.ProductId, order.OrderStatus); err != nil {
		return err
	}
	return
}

//查询一条数据
func (o *OrderManager) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return &datamodels.Order{}, err
	}
	selectOneSql := "select * from " + o.table + " where ID=" + strconv.FormatInt(orderID, 10)
	row, err := o.mysqlConn.Query(selectOneSql)
	if err != nil {
		return &datamodels.Order{}, err
	}
	result := common.GetResultRow(row)
	order = &datamodels.Order{}
	if len(result) == 0 {
		return &datamodels.Order{}, nil
	}
	//数据映射到结构体
	common.DataToStructByTagSql(result, order)
	return
}

//查询所有数据
func (o *OrderManager) SelectAll() (orderArray []*datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}
	selectAllSql := "select * from " + o.table
	rows, err := o.mysqlConn.Query(selectAllSql)
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	//数据映射到结构体
	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orderArray = append(orderArray, order)
	}
	return
}

//查询订单相关
func (o *OrderManager) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}
	selectAllSql := "select o.ID,p.ProductName,o.OrderStatus from " + o.table + "as o left join product p on o.productID = p.productID"
	rows, err := o.mysqlConn.Query(selectAllSql)
	if err != nil {
		return nil, err
	}
	return common.GetResultRows(rows), err

}
