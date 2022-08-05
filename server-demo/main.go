package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"weyii.co/comm/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://agriPP-dev:ql2k22GO@weyii.co:27017/agriPP?authSource=agriPP"

var client *mongo.Client

var db map[string]string = map[string]string{
	"foo":   "bar", // user:foo password:bar
	"manu":  "123", // user:manu password:123
	"admin": "123456",
	"org1":  "111111",
	"org2":  "222222",
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	user := r.Group("/user")

	user.POST("get", func(c *gin.Context) {
		//param := make(map[string]interface{})
		param := Param{}
		c.ShouldBindJSON(&param)
		log.Printf("%v", &param)

		// 查询多个
		// 将选项传递给Find()
		findOptions := options.Find()
		findOptions.SetLimit(20)

		// 定义一个切片用来存储查询结果
		var results []*Procurement

		// 把bson.D{{}}作为一个filter来匹配所有文档
		//cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		//if err != nil {
		//	log.Fatal(err)
		//}

		collection := client.Database("govProcurement").Collection("project")

		/*
		 以下证明：可以使用bson.M{}作为 filter，也可以使用[]bson.E{}作为 filter
		*/
		//filter := map[string]interface{}{}

		/*
			//way 1: 使用 bson.M
			filter := bson.M{}

			filter["status"] = param.Status[0]
			filter["pubDate"] = []bson.E{
				bson.E{Key: "$gte", Value: time.Unix(param.STime, 0).Format("2006-01-02")},
				bson.E{Key: "$lte", Value: time.Unix(param.ETime, 0).Format("2006-01-02")},
			}
		*/

		//way 2: 使用 []bson.E 作为 filter
		filter := []bson.E{}

		filter = append(filter, bson.E{"status", param.Status[0]})
		filter = append(filter, bson.E{"pubDate", []bson.E{
			bson.E{Key: "$gte", Value: time.Unix(param.STime, 0).Format("2006-01-02")},
			bson.E{Key: "$lte", Value: time.Unix(param.ETime, 0).Format("2006-01-02")},
		}})
		log.Printf("filter: %q", filter)

		//findFilter, _ := bson.Marshal(filter)

		/*
			//可以直接将 filter 作为Find()查询条件
			cur, _ := collection.Find(context.Background(), filter, options.Find())
		*/

		//也可使用[]bson.D 包装之后，作为 Aggregte()的聚合条件(pipeline)。此时需注意：$sort 值一定要使用 bson.D
		cur, err := collection.Aggregate(
			context.Background(),
			[]bson.D{
				bson.D{{"$match", filter}},
				bson.D{{"$sort", bson.D{{"pubDate", -1}}}},
			},
			options.Aggregate().SetMaxTime(4*time.Second))

		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.Background())

		if err = cur.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, results)
	})

	//用户动态验证码请求
	user.GET("/validCodeReq"), func(c *gin.Context) {
		//generate valid code

		//generate valid code image

		//generate requestId

		//store requestId in Redis

		//return imgage and requestId
	}

	//用户注册
	user.POST("/register", func(c *gin.Context) {
		//TODO get registed param
		username	:= c.PostForm("username")
		passwd		:= c.PostForm("passwd")
		requestId	:= c.PostForm("requestId")
		validCode	:= c.PostForm("validCode")
		//passwd		:= c.DefaultPostForm("passwd", "0.1")

		//sha256(password + salt)

		//insert MongoDB

		//return user regist result

	})

	r.Use(gin.Recovery())
	return r
}

func main() {
	r := setupRouter()
	//flag.StringVar(&common.ServeStationId, "s", "A", "station of flag")
	//flag.IntVar(&common.Port, "p", 8080, "port of running")
	//flag.Parse()
	//fmt.Println(common.Port)

	//init mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Listenand Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%v", 8080))
}
