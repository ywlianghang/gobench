package CURD

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type writeJobStruct struct {
	writech chan int
	total int
	n int
}

func InsertTable(job writeJobStruct,tablename string,tableSize int,filenum string){
    buf := HtgoodsInsert(job.total,tablename,filenum)
	for i:=0; i<=job.total; i++{
		value := HtgoodsValue(filenum)
		if i == job.total{
			buf = append(buf, []byte(value)...)
			buf = append(buf,";"...)
		}else {
			buf = append(buf, []byte(value)...)
			buf = append(buf,","...)
		}
		time.Sleep(10 * time.Millisecond)
	}
	ss :=string(buf)
	fmt.Printf("第" + strconv.Itoa(job.n) + "次插入%d条数据！\n",tableSize)
	_,err := DB.Exec(ss)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("完成---" + strconv.Itoa(job.n) + "次插入%d条数据！\n",tableSize)
	job.writech<- 1
}


