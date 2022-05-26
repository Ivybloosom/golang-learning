# 说明

### 准备数据库连接驱动
如果没有下载过数据库驱动，先运行如下命令安装 mysql 驱动
```bash
go get -u github.com/go-sql-driver/mysql
```

### 准备数据表
```sql
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '姓名',
  `phone` int(12) DEFAULT '0' COMMENT '电话',
  `email` varchar(50) DEFAULT '0' COMMENT '邮箱',
  `addTime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';
```