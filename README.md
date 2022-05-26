# notes of golang

---

### go get报错 i/o timeout的解决办法

GOPROXY

Go1.11新增的环境变量，如果设置了此环境变量，那么在下载依赖时，会从环境变量设置的代理地址去下载。开源项目goproxyio可以帮助开发者一键构建自己的代理服务。并且提供了一个公用的代理服务https://goproxy.io。设置方法如下：

```bash
// 启用 Go Modules 功能
export GO111MODULE=on
// 配置 GOPROXY 环境变量
export GOPROXY="https://goproxy.io"
```

