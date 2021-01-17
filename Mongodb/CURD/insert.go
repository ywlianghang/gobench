package CURD

//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//	"time"
//)
//
//var MongoClient *mongo.Client
//
//func mongodb() {
//	param := fmt.Sprintf("mongodb://XXX.XXX.XXX.XXX:27017")
//	clientOptions := options.Client().ApplyURI(param)
//
//	// 建立客户端连接
//	client, err := mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//		fmt.Println(err)
//	}
//
//	// 检查连接情况
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//		fmt.Println(err)
//	}
//	fmt.Println("Connected to MongoDB!")
//
//	//指定要操作的数据集
//	collection := client.Database("ccmsensor").Collection("mtr")
//
//	//执行增删改查操作
//	type CurlInfo struct {
//		DNS float64 `json:"NAMELOOKUP_TIME"` //NAMELOOKUP_TIME
//		TCP float64 `json:"CONNECT_TIME"`    //CONNECT_TIME - DNS
//		SSL float64 `json:"APPCONNECT_TIME"` //APPCONNECT_TIME - CONNECT_TIME
//	}
//
//	// 新增
//	func insertSensor(client *mongo.Client, collection *mongo.Collection) (insertID primitive.ObjectID) {
//		apps := make(map[string]ConnectData, 0)
//		apps["app1"] = ConnectData{
//			Latency:  30.983999967575,
//			RespCode: 200,
//			Url:      "",
//			Detail: CurlInfo{
//				DNS: 5.983999967575,
//				TCP: 10.983999967575,
//				SSL: 15.983999967575,
//			},
//		}
//
//		record := &Sensor{
//			Clientutc: time.Now().UTC().Unix(),
//			ISP:       "China Mobile",
//			DataByAPP: apps,
//		}
//
//		insertRest, err := collection.InsertOne(context.TODO(), record)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		insertID = insertRest.InsertedID.(primitive.ObjectID)
//		return insertID
//	}
//	//查询
//	timestamp := time.Now().UTC().Unix()
//	start := timestamp - 180
//	end := timestamp
//
//	filter := bson.D{
//		{"isp", isp},
//		{"$and", bson.A{
//			bson.D{{"clientutc", bson.M{"$gte": start}}},
//			bson.D{{"clientutc", bson.M{"$lte": end}}},
//		}},
//	}
//
//
//	// 断开客户端连接
//	err = client.Disconnect(context.TODO())
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connection to MongoDB closed.")
//}
