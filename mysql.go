package main

import (
	"learn/mysql"
)

func main() {
	//pl := mysql.PlainSQL{}
	//pl.SqlStatement()
	orm := mysql.ORM{UserName: "root", PassWord: "root", DbName : "test_db"}
	orm.Migration()

}

