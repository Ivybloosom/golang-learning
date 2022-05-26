# 运行方式 1: 使用软件 netcat
命令行:
```shell
./clock &
nc localhost 8080
```
```shell
nc localhost 8080
```
```shell
nc localhost 8080
```
...
停止方式：
```shell
^C
killall clock
```


# 运行方式 2: 使用go - net.Dial
命令行:	
```shell
./clock &
./netcat1
```
```shell
./netcat1
```
```shell
./netcat1
```
...

停止方式： 
```shell
^C
killall clock
```