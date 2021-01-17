package main

import (
	"awesomeProject/gobench/mysql/CURD"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	CURD.Mysql()
}
