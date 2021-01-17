package CURD

import (
	"awesomeProject/gobench/readom"
	"fmt"
)

type readJobStruct struct {
	readch chan int
	total int
	n int
}

func SelectTable(job readJobStruct,tablename string) {
	 tenId := string(readom.RandomInt())
   sql := HtgoodsSelect(job.total,tablename)
   rows,err := DB.Query(string(sql),tenId)
   if err != nil {
   	fmt.Printf("query failed,err:%v\n",err)
		 return
	 }
	 for rows.Next(){
	 	var uid int
	 	err = rows.Scan(&uid)
	 	fmt.Println(uid)
	 }
	 job.readch <- 1

}