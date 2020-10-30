package common

import "database/sql"

// 创建MySQL连接
func NewMySQLConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:chen0423@tcp(39.106.88.184:3306)/product?charset=utf8")
	return
}

//获取一条数据
func GetResultRow(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	records := make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				records[columns[i]] = string(v)
			}
		}

	}
	return records
}

//获取所有数据
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	//返回所有列
	columns, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	values := make([][]byte, len(columns))
	//这里表示一行填充数据
	scans := make([]interface{}, len(columns))
	//这里scans引用values，把数据填充到[]byte里
	for k, _ := range values {
		scans[k] = &values[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把values中的数据复制到row中
		for k, v := range values {
			key := columns[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	return result
}
