package repositories

import (
	"database/sql"
	"spike/common"
	"spike/datamodels"
	"strconv"
)

type IUserRepository interface {
	Conn() error
	Select(userName string) (user *datamodels.User, err error)
	Insert(user *datamodels.User) (userId int64, err error)
}

type UserManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func NewUserManagerRepository(table string, db *sql.DB) IUserRepository {
	return &UserManagerRepository{table, db}
}

func (u *UserManagerRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql, err := common.NewMySQLConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u *UserManagerRepository) Select(userName string) (user *datamodels.User, err error) {
	if userName == "" {
		return &datamodels.User{}, nil
	}
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, nil
	}
	selectSql := "select * from " + u.table + " where userName=?"
	row, err := u.mysqlConn.Query(selectSql, userName)
	if err != nil {
		return &datamodels.User{}, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.User{}, nil
	}
	//数据映射到结构体
	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return

}

func (u *UserManagerRepository) Insert(user *datamodels.User) (userId int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	insertSql := "insert into " + u.table + "(nickName,  userName, password) values(?,?,?)"
	stmt, err := u.mysqlConn.Prepare(insertSql)
	if err != nil {
		return userId, err
	}
	result, err := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if err != nil {
		return userId, err
	}
	return result.LastInsertId()

}

func (u *UserManagerRepository) SelectById(userId int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	selectOneSql := "select * from " + u.table + " where id =" + strconv.FormatInt(userId, 10)
	row, err := u.mysqlConn.Query(selectOneSql)
	if err != nil {
		return &datamodels.User{}, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.User{}, err
	}
	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return
}
