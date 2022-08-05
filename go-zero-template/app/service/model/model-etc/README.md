## 运行 model.go 文件(需要安装goctl工具)
```bash
# 生成 model 代码
goctl model mongo -type User -dir . #mongo不带cache
goctl model mongo -type User -dir . -cache #mongo带cache
```