package test

import (
	"apiproject/util"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"log"
	"testing"
	"time"
)
import "github.com/globalsign/mgo"

type User struct {
	Id             bson.ObjectId `bson:"_id"`
	Username       string
	Interests      []string
	CreateTime     time.Time `bson:"createTime"`
	CreateTimeUnix int64     `bson:"createTimeUnix"`
	CreateTimeStr  string    `bson:"createTimeStr"`
}

var collection *mgo.Collection

func init() {
	url := "mongodb://localhost:27017/test_collect"
	// 解析MongoDB参数
	dialInfo, err := mgo.ParseURL(url)
	// 1、连接MongoDB
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	// 2、选择数据库
	db := session.DB(dialInfo.Database)

	// 3、选择表（集合）
	collection = db.C("table1")
}

/**
新增
*/
func TestInsert(t *testing.T) {
	//新增数据
	stu1 := new(User)
	stu1.Id = bson.NewObjectId()
	stu1.Username = "stu1_name"
	stu1.Interests = []string{"象棋", "游泳", "跑步"}
	err := collection.Insert(stu1)
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}
}

/**
新增多条
*/
func TestInsertMany(t *testing.T) {
	createTime := time.Now()
	createTimeUnix := createTime.Unix()
	createTimeStr := util.FormatTime(createTime)

	//新增数据
	stu2 := User{bson.NewObjectId(), "stu2_name", []string{"象棋", "游泳", "跑步"}, createTime, createTimeUnix, createTimeStr}
	stu3 := User{bson.NewObjectId(), "stu3_name", []string{"象棋", "游泳", "跑步"}, createTime, createTimeUnix, createTimeStr}
	stu4 := User{bson.NewObjectId(), "stu4_name", []string{"象棋", "游泳", "跑步"}, createTime, createTimeUnix, createTimeStr}
	userList := []interface{}{}
	userList = append(userList, stu2, stu3, stu4)

	err := collection.Insert(userList...)
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}
}

/**
修改
*/
func TestUpdate(t *testing.T) {
	interests := []string{"象棋2", "游泳2", "跑步"}
	err := collection.Update(bson.M{"_id": bson.ObjectIdHex("5d42a43d5d5cb277f2925435")}, bson.M{"$set": bson.M{
		"interests": interests,
	}})
	if err != nil {
		fmt.Println("修改失败")
	} else {
		fmt.Println("修改成功")
	}
}

/**
删除
*/
func TestDelete(t *testing.T) {
	err := collection.Remove(bson.M{"_id": bson.ObjectIdHex("5d42a43d5d5cb277f2925435")})
	if err != nil {
		fmt.Println("删除失败" + err.Error())
	} else {
		fmt.Println("删除成功")
	}
}

/**
查询单个记录
*/
func TestFindOne(t *testing.T) {
	//根据ObjectId进行查询
	user := new(User)
	collection.Find(bson.M{"_id": bson.ObjectIdHex("5d42a2290e84057cfa547c9c")}).One(user)
	fmt.Println(user)
}

/**
查询多条记录
*/
func TestFindMany(t *testing.T) {
	//    defer session.Close()
	var users []User
	//    c.Find(nil).All(&users)
	collection.Find(bson.M{"username": "stu2_name"}).All(&users)
	log.Println(users)
}

/**
测试mongo的默认UTC时区对日期时间的应用的影响
*/
func TestMongoTime(t *testing.T) {
	//新增数据
	stu1 := new(User)
	stu1.Id = bson.NewObjectId()
	stu1.Username = "stu1_name"
	stu1.Interests = []string{"象棋", "游泳", "跑步"}
	stu1.CreateTime = time.Now()
	stu1.CreateTimeUnix = stu1.CreateTime.Unix()
	stu1.CreateTimeStr = util.FormatTime(stu1.CreateTime)
	err := collection.Insert(stu1)
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}

	//读取
	var users []User
	collection.Find(nil).All(&users)
	for _, user := range users {
		log.Println(user.CreateTime.Unix())
		log.Println(user.CreateTimeUnix)
		now := time.Now()
		log.Println("当前时间:", now)
		log.Println("用户创建时间:", user.CreateTime, user.CreateTimeStr)
		log.Println("当前时间-用户创建时间", now.Sub(user.CreateTime))

		//修改时间入库
		timeNew := user.CreateTime.Add(10 * time.Second)
		log.Println("timeNew", timeNew, timeNew.Unix())
		err := collection.Update(bson.M{"_id": user.Id}, bson.M{"$set": bson.M{
			"createTime":     timeNew,
			"createTimeUnix": timeNew.Unix(),
			"createTimeStr":  util.FormatTime(timeNew),
		}})
		if err != nil {
			fmt.Println("修改失败")
		} else {
			fmt.Println("修改成功")
			log.Println("用户修改后入库前的createTimeStr", util.FormatTime(timeNew))
		}

		//再次查询
		var user2 User
		collection.Find(bson.M{"_id": user.Id}).One(&user2)
		now2 := time.Now()
		log.Println("当前时间:", now2)
		log.Println("用户创建时间:", user2.CreateTime, user2.CreateTimeStr)
		log.Println("当前时间-用户创建时间", now2.Sub(user2.CreateTime))
	}
	log.Println(users)
}
