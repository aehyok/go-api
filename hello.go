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
	fmt.Printf("insert success\n")
	addRecord(Db)
	updateRecord(Db)
	deleteRecord(Db)
}

func addRecord(Db *sqlx.DB) {
	for i := 0; i < 2; i++ {
		result, err := Db.Exec("insert into user (id,account,nickName,password) values(5,'111', '222', '3333')")
		if err != nil {
			fmt.Printf("data insert faied, error:[%v]", err.Error())
			return
		}
		id, _ := result.LastInsertId()
		fmt.Printf("insert success, last id:[%d]\n", id)
	}
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
