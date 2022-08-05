# goctl 生成代码需要的api配置文件

api 文件语法介绍： https://go-zero.dev/cn/docs/design/grammar

## 运行 api 文件(需要安装goctl工具)

```bash
# 创建 api 服务
goctl api new xxx  #xxx:模块名
goctl api go -api xxx.api -dir .
```