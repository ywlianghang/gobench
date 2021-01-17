package CURD

import (
	"awesomeProject/gobench/flage"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

var DB *sql.DB
func writeChan(connInfo map[string]string){
	dbname := connInfo["database"]
	tableNum, _ := strconv.Atoi(connInfo["table"])
	filenum := connInfo["filenum"]
	tablename := connInfo["tablename"]
	for i := 1; i <= tableNum; i++ {
		tablename := tablename + strconv.Itoa(i)
		if IsexitTable(dbname, tablename) == false {
			CreateTable(dbname, tablename,filenum)
		}
		fmt.Println("====insert start=====")
		start := time.Now()
		total, _ := strconv.Atoi(connInfo["tableSize"])
		gonum, _ := strconv.Atoi(connInfo["threads"])

		// 测试插入数据库的功能，每次最多同步20个工作协程
		writejobChan := make(chan writeJobStruct, gonum)
		go writeWorker(writejobChan, tablename, total, filenum)
		// 统计使用次数
		writech := make(chan int, gonum)

		for n := 0; n < gonum; n++ {
			writejob := writeJobStruct{
				writech:    writech,
				total: total,
				n: n,
			}
			writejobChan <- writejob
		}
		ii := 0
		for {
			<-writech
			ii++
			if (ii >= gonum) {
				break;
			}
		}
		end := time.Now()
		curr := end.Sub(start)
		fmt.Println("====insert end=====")
		fmt.Println("insert run time:", curr)
	}
}

func readChan(connInfo map[string]string){
	tableNum, _ := strconv.Atoi(connInfo["table"])
	tablename := connInfo["tablename"]

	for i := 1; i <= tableNum; i++ {
		tablename := tablename + strconv.Itoa(i)
		fmt.Println("====query start=====")
		start := time.Now()
		total, _ := strconv.Atoi(connInfo["tableSize"])
		gonum, _ := strconv.Atoi(connInfo["threads"])

		// 测试插入数据库的功能，每次最多同步20个工作协程
		readjobChan := make(chan readJobStruct,gonum)
		go readWorker(readjobChan, tablename)
		// 统计使用次数
		readch := make(chan int, gonum)

		for n := 0; n < gonum; n++ {
			readjob := readJobStruct{
				readch:    readch,
				total: total,
				n : n,
			}
			readjobChan <- readjob
		}

		ii := 0
		for {
			<-readch
			ii++
			if (ii >= gonum) {
				break;
			}
		}
		end := time.Now()
		fmt.Println("====query end=====")
		curr := end.Sub(start)
		fmt.Println("query run time:", curr)
	}
}

func Mysql() {
	connInfo := flage.Help()
	DSN := Dsnn(connInfo)
	read_write_mode := connInfo["read_write_mode"]
	conn, err := sql.Open("mysql", DSN)

	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		log.Fatal("Ping mysql failed, err:%s", err.Error()) //proper error handling instead
	}
	conn.SetConnMaxLifetime(time.Second * 500) //设置连接超时500秒
	conn.SetMaxIdleConns(50)
	conn.SetMaxOpenConns(1000) //设置最大连接数
	DB = conn
	if read_write_mode == "write_only"{
		writeChan(connInfo)
	}
	if read_write_mode == "read_only"{
		readChan(connInfo)
	}
	if read_write_mode == "read_write" {
		writeChan(connInfo)
		readChan(connInfo)
	}

}
