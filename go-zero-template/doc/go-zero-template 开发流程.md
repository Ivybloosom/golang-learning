# go-zero-template 开发流程
---
[TOC]
---
## 1. 生成代码
### 1.1 REST API 部分
#### 1.1.1 编写 api 文件
#### 1.1.2. 运行 api 文件(需要安装goctl工具)
```bash
# 创建 api 服务
goctl api new xxx  #xxx:模块名
goctl api go -api xxx.api -dir .
```
### 1.2 数据实体 model 部分
#### 1.2.1 编写 model.go 文件
#### 1.2.2 运行 model.go 文件(需要安装goctl工具)
```bash
# 生成 model 代码
goctl model mongo -type User -dir . #mongo不带cache
goctl model mongo -type User -dir . -cache #mongo带cache
```
## 2. 调整服务配置，连接中间件
### 2.1 修改 yaml 配置文件
```yaml
# mongo配置
Mongodb:
  Uri: "mongodb://localhost:27017"
  Db: "test"
  MaxPoolSize: 2000
  
# jwt配置
Auth:
  AccessSecret: 63AB37EA-748C-1E44-E268-B778B3DB616F
  AccessExpire: 259200 #72h
```

### 2.2 修改配置文件 Config, 和 yaml 文件对齐
```go
type Config struct {
	// rest api配置
	rest.RestConf
	// mongo数据库配置
	Mongodb MongoConf
	// jwt配置
	Auth AuthConfig
}
type MongoConf struct {
	Uri         string
	Db          string
	MaxPoolSize uint64
}
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
```

### 2.3 填充服务依赖 svc
## 3. 填充业务逻辑 logic