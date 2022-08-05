# go-zero-template 运行方式
---
[TOC]
---
## 1.配置环境
### 1.1 修改`app/service/cmd/api/etc/example-api.yaml`文件
```yaml
# mongo配置
Mongodb:
  Uri: $your_mongo_path
  Db: $you_database_name
```
示例: 本地运行的mongo，连接其中的名字为`test`的数据库
```yaml
# mongo配置
Mongodb:
  Uri: "mongodb://localhost:27017"
  Db: "test"
```

## 2.运行示例api服务

```bash
cd app/service/cmd/api/
go run example.go
```
服务启动响应
```bash
mongo conn success!
Starting server at 0.0.0.0:8888...
```

## 3. 测试调用接口 

注意 : 这里调用的路径中的`$your_user_id`是你自己数据库中的一个id，返回的数据中也是你自己数据库中的数据

GET 方法 `http://localhost:8888/user/getuser/$your_user_id`

http响应
```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "ObjectID(\"XXXXXXXXXXXXXX\")",
        "name": "XXX"
    }
}
```