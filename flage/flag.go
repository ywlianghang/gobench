package flage

import (
	"flag"
	"fmt"
	"os"
)

func Help() map[string]string{
	/*
		定义变量接收控制台参数
	*/
	var username string
	var password string
	var host string
	var port string
	var database string
	var charset string
	var threads string
	var tableSize string
    var table string
	var connInfo map[string]string
	var filenum string
	var read_write_mode string
	var tablename string
	connInfo = make(map[string]string,20)

	// 添加 --help帮助信息
	flag.StringVar(&username,"username","","用户名，默认为空")
	flag.StringVar(&password,"password","","密码，默认为空")
	flag.StringVar(&host,"host","127.0.0.1","主机名，默认 127.0.0.1")
	flag.StringVar(&port,"port","3306","端口号，默认为空")
	flag.StringVar(&database,"database","test","数据库名称，默认为test")
	flag.StringVar(&charset,"charset","utf8mb4 ","连接数据库的字符集，默认 utf8mb4")
	flag.StringVar(&tablename,"tablename","gobench","表名")
	flag.StringVar(&threads,"threads","1","并发执行数")
    flag.StringVar(&tableSize,"tableSize","100","单个并发插入的数据量")
    flag.StringVar(&table,"table","1","创建几个表")
	flag.StringVar(&filenum,"filenum","10","每个表需要多少个列，默认8个，超过则使用file1、file2标注")
	flag.StringVar(&read_write_mode,"read_write_mode","write_only","配置参数有read_only,write_only,read_write")
	flag.Parse()
	connInfo["username"] = username
	connInfo["password"] = password
	connInfo["host"] = host
	connInfo["port"] = port
	connInfo["database"] = database
	connInfo["charset"] = charset
	connInfo["tablename"] = tablename
	connInfo["threads"] = threads
	connInfo["tableSize"] = tableSize
	connInfo["table"] = table
	connInfo["filenum"] = filenum
	connInfo["read_write_mode"] = read_write_mode

	if username == "" || password == ""{
		fmt.Println("当前用户名或密码为空，请正确输入！")
		os.Exit(0)
	}
	return connInfo
}
