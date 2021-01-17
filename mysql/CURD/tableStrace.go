package CURD

import (
	"awesomeProject/gobench/readom"
	"strconv"
	"time"
)

func TableStract(dbname string,tablename string,filenum string) string{
		files,_ :=strconv.Atoi(filenum)
		var tableStructure []byte
		tableStructure = append(tableStructure,"CREATE TABLE "...)
		tableStructure = append(tableStructure,"`"...)
		tableStructure = append(tableStructure,dbname...)
	    tableStructure = append(tableStructure,"`"...)
	    tableStructure = append(tableStructure,"."...)
	    tableStructure = append(tableStructure,"`"...)
	    tableStructure = append(tableStructure,tablename...)
	    tableStructure = append(tableStructure,"`"...)
		tableStructure = append(tableStructure," (`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',"...)
		tableStructure = append(tableStructure,"`tenantry_id` bigint(20) NOT NULL COMMENT '商品id',"...)
	    tableStructure = append(tableStructure,"`code` varchar(64) NOT NULL COMMENT '商品编码（货号）',"...)
		tableStructure = append(tableStructure,"`goods_name` varchar(50) NOT NULL COMMENT '商品名称',"...)
		tableStructure = append(tableStructure,"`props_name` varchar(100) NOT NULL COMMENT '商品名称描述字符串，格式：p1:v1;p2:v2，例如：品牌:盈讯;型号:F908',"...)
		tableStructure = append(tableStructure,"`price` decimal(10,2) NOT NULL COMMENT '商品定价',"...)
	    tableStructure = append(tableStructure,"`price_url` varchar(1000) NOT NULL COMMENT '商品主图片地址',"...)
		if files > 10{
			for i := 0; i < files-10; i++{
				tableStructure = append(tableStructure,"`files"...)
				tableStructure = append(tableStructure,strconv.Itoa(i)...)
				tableStructure = append(tableStructure,"` varchar(50) NOT NULL COMMENT '商品列',"...)
			}
		}
		tableStructure = append(tableStructure,"`create_time` datetime NOT NULL COMMENT '商品创建时间',"...)
		tableStructure = append(tableStructure,"`modify_time` datetime DEFAULT NULL COMMENT '商品最近修改时间',"...)
		tableStructure = append(tableStructure,"`deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '标记逻辑删除',"...)
		tableStructure = append(tableStructure,"PRIMARY KEY (`id`)"...)
		tableStructure = append(tableStructure,") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品信息表';"...)
		return string(tableStructure)
}


func HtgoodsInsert(job int,tablename string,filenum string) []byte{
	var buf []byte
	files,_ :=strconv.Atoi(filenum)
	buf = append(buf,"INSERT INTO "...)
	buf = append(buf,[]byte(tablename)...)
	buf = append(buf," (tenantry_id,code,goods_name,props_name,price,price_url,"...)
	if files > 10{
		for i := 0; i < files-10; i++{
			buf = append(buf,"files" + strconv.Itoa(i)...)
			buf = append(buf,","...)
		}
	}
	buf = append(buf,"create_time,modify_time) values"...)
    return buf
}
func HtgoodsValue(filenum string) string{
	var val []byte
    var props_name []byte
	files,_ :=strconv.Atoi(filenum)

	tenId := readom.RandomInt()
	code := readom.RandomString()
	goods_name := readom.RandomName()

	props_name = append(props_name,[]byte(goods_name)...)
	props_name = append(props_name,":"...)
	props_name = append(props_name,[]byte(code)...)
	pice := readom.RandomFloat()
	pic_url := readom.RandUrl()
	createTime := readom.RandomDate()
	modifyTime := time.Now().Format("2006-01-02 15:04:05")
	//title := fmt.Sprintf("%s DIOR香水三件套小样花果香调（真我+魅惑+花漾各5ml",props_name)

    val = append(val," ("...)
    val = append(val,tenId...)
    val = append(val,","...)
    val = append(val,"'"...)
    val = append(val,code...)
	val = append(val,"'"...)
	val = append(val,","...)
	val = append(val,"'"...)
    val = append(val,goods_name...)
	val = append(val,"'"...)
	val = append(val,","...)
	val = append(val,"'"...)
	val = append(val,props_name...)
	val = append(val,"'"...)
	val = append(val,","...)
	val = append(val,pice...)
	val = append(val,","...)
	val = append(val,"'"...)
	val = append(val,pic_url...)
	val = append(val,"'"...)
	val = append(val,","...)
	if files > 10{
		for i := 0; i < files-10; i++{
			val = append(val,"'"...)
			val = append(val,code...)
			val = append(val,"'"...)
			val = append(val,","...)
		}
	}
	val = append(val,"'"...)
	val = append(val,createTime...)
	val = append(val,"'"...)
	val = append(val,","...)
	val = append(val,"'"...)
	val = append(val,modifyTime...)
	val = append(val,"'"...)
	val = append(val,")"...)
    return string(val)
}

func HtgoodsSelect(job int,tablename string) []byte{
	buf := make([]byte,0,job)
	buf = append(buf,"SELECT * from  "...)
	buf = append(buf,[]byte(tablename)...)
	buf = append(buf," where tenantry_id =  "...)
	buf = append(buf,"?"...)
	return buf
}