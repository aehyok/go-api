package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "Hk1997mc1999!!!"
	ipAddrees string = "139.186.205.7"
	port      int    = 3306
	dbName    string = "aehyok"
	charset   string = "utf8"
)

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func updateRecord(Db *sqlx.DB) {
	//更新uid=1的username
	result, err := Db.Exec("update user set nickName = '222222' where account = '111'")
	if err != nil {
		fmt.Printf("update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

func main() {
	var Db *sqlx.DB = connectMysql()
	defer Db.Close()
	selectRecord(Db)
	fmt.Printf("select success\n")
	// addRecord(Db)
	// updateRecord(Db)
	// deleteRecord(Db)
}

func addRecord(Db *sqlx.DB) {
	for i := 0; i < 2; i++ {
		result, err := Db.Exec("insert into user (id,account,nickName,password) values(7,'111', '222', '3333')")
		if err != nil {
			fmt.Printf("data insert faied, error:[%v]", err.Error())
			return
		}
		id, _ := result.LastInsertId()
		fmt.Printf("insert success, last id:[%d]\n", id)
	}
}

func selectRecord(Db *sqlx.DB) {
	rows, err := Db.Query(" select * from  user")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//列出所有查询结果的字段名
	cols, _ := rows.Columns()

	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	// res := make([]map[string]string, 0)
	for rows.Next() {
		_ = rows.Scan(scans...)
		row := make(map[string]string)
		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		fmt.Println(row)
		// res = append(res, row)
	}
	// fmt.Println(res)
}

func deleteRecord(Db *sqlx.DB) {
	//删除uid=2的数据
	result, err := Db.Exec("delete from user where account = '111'")
	if err != nil {
		fmt.Printf("delete faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("delete success, affected rows:[%d]\n", num)
}
