package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	// rest api配置
	rest.RestConf
	// mongo数据库配置
	Mongodb MongoConf
}
type MongoConf struct {
	Uri         string
	Db          string
	MaxPoolSize uint64
}
