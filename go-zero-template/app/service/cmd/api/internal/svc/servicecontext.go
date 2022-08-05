package svc

import (
	"context"
	"fmt"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/config"
	"github.com/Ivybloosom/go-zero-template/app/service/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var ServiceContextObj *ServiceContext

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	RedisClient   *redis.Redis
	MongoClient   *mongo.Client
	MongoClientDb *mongo.Database
}

func NewServiceContext(c config.Config) *ServiceContext {
	srvCtx := &ServiceContext{
		Config: c,
	}
	////redis初始化
	//if c.Redis.Host != "" {
	//	srvCtx.RedisClient = srvCtx.Config.Redis.NewRedis()
	//}
	//mongo初始化
	if c.Mongodb.Uri != "" {
		client, db, err := srvCtx.initMongoDB(c)
		if err != nil {
			panic("Mongodb init error" + err.Error())
		}
		srvCtx.MongoClient = client
		srvCtx.MongoClientDb = db
	}
	ServiceContextObj = srvCtx
	return srvCtx
}

// 初始化MongoDB
func (s *ServiceContext) initMongoDB(c config.Config) (*mongo.Client, *mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//设置mongo参数
	options := options.Client().
		ApplyURI(c.Mongodb.Uri).
		SetMaxPoolSize(c.Mongodb.MaxPoolSize)
	//常见数据库连接
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("mongo conn success!")
	pref := readpref.ReadPref{}
	err = client.Ping(ctx, &pref)
	db := client.Database(c.Mongodb.Db)
	if err != nil {
		return nil, nil, err
	}
	return client, db, nil
}
