package CURD

import (
	"fmt"
	"log"
)

func Dsnn(conninfo map[string]string) string{
	connString := conninfo["username"] + ":" + conninfo["password"] + "@" +"tcp(" + conninfo["host"] + ":" +conninfo["port"] + ")/" + conninfo["database"] + "?charset=" + conninfo["charset"]
	return connString
}

func IsexitTable(dbname string,tablename string) bool{
	var Database string
	status := false
	sql := fmt.Sprintf("show tables from %s",dbname)
	row ,err := DB.Query(sql)
	if err != nil{
		log.Fatal("query table sql fail,error info ",err)
	}
	for row.Next() {
		err = row.Scan(&Database)
		if Database == tablename {
			status = true
		}
	}
	return status
}

func CreateTable(dbname string,tablename string,filenum string) {
	tableStructure := TableStract(dbname,tablename,filenum)
	stmt,err := DB.Prepare(tableStructure)
	_,err = stmt.Exec()
	if err != nil{
		fmt.Printf("创建 %s 库下%s 表失败,失败原因：%s\n",dbname,tablename,err)
	}else {
		fmt.Printf("创建 %s 库下%s 表成功。\n",dbname,tablename)
	}
}

func writeWorker(writejobChan <-chan writeJobStruct,tablename string,tableSize int,filenum string){
		for job := range writejobChan {
			go InsertTable(job, tablename,tableSize,filenum)
		}
}

func readWorker(readjobChan <-chan readJobStruct,tablename string){
	for job := range readjobChan {
		go SelectTable(job,tablename)
	}
}