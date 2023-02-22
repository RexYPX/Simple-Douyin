# Simple-Douyin
Simple Douyin is a project imitated from ByteDance's Douyin（Tiktok）

# Feature
1. 采用开源 HTTP 框架 Hertz、开源 RPC 框架（Kitex）及开源 ORM 框架 GORM 开发，基于 RPC 微服务 + Hertz 提供 HTTP 服务 + GORM 实现 ORM

2. 基于《接口文档在线分享- Apifox》提供的接口进行开发，使用《极简抖音App使用说明 - 青训营版 》提供的APK进行Demo测试， 功能完整实现 ，前端接口匹配良好。

3. 代码结构采用 (HTTP API 层 + RPC Service 层 + Dal 层) 项目结构清晰 ，代码符合规范

4. 使用 JWT 进行用户token的校验

5. 使用 ETCD 进行服务发现和服务注册；

6. 使用 Gorm 对 MySQL 进行 ORM 操作；

7. 使用 Hertz 中间件 tracer 实现链路跟踪；

8. 使用 Pprof 中间件实现代码性能检测

9. 使用 Redis 中间件作为数据库缓存

# Environment
- go1.19.5 linux/amd64
- hertz v0.5.2
- kitex v0.4.4

# Quick Start
1. Edit pkg/constant.go to config your project

2. Setup basic dependency(please install docker & docker-compose previously)
```shell
make start
```

3. run user service
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

4. run feed service
```shell
cd cmd/feed
sh build.sh
sh output/bootstrap.sh
```

5. run publish service
```shell
cd cmd/publish
sh build.sh
sh output/bootstrap.sh
```

6. run comment service
```shell
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

7. run favorite service
```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

8. run relation service
```shell
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

9. run message service
```shell
cd cmd/message
sh build.sh
sh output/bootstrap.sh
```

10. run api service
```shell
cd cmd/api
go build
./api
```
publishList获取视频列表流程图：
![publishList](./pic/publishList.jpg)


favoriteAction点赞流程图：
![favoriteAction](./pic/点赞.jpg)


favoriteList获取点赞列表流程图：
![favoriteList](./pic/点赞列表.jpg)
