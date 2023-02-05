# 20230205
GQY:
1. 添加 cmd/api/biz/handler/pack/comment.go
2. 修改 cmd/api/biz/handler/comment_service.go 和 cmd/api/rpc/comment.go，删除 handler.go ，并完成API验证

需要配合完成的功能：在 cmd/comment/test 中的接口


# 20230204
YPX：
relation模块后端处理逻辑及简单RPC测试完成：
1. 编写 cmd/relation 下代码，实现用户关注/取关、返回关注列表、粉丝列表、好友列表功能（已完成SQL验证）
2. 修改 pkg/constants/constant.go 代码，添加 "RelationTableName" 字段

仍需配合完成功能：
1. cmd/relation/rpc/user.go 中，需要 user 模块提供 Token2Id 及 Id2User 两个接口

GQY：
comment模块后端处理逻辑及简单RPC测试完成：
1. 编写 cmd/comment 下代码，实现用户增加/删除评论，根据视频编号返回评论列表（已完成SQL验证）
2. 修改 pkg/constants/constant.go ，添加相关常量表示
3. 编写 cmd/api/biz/handler/comment_service.go 和 handler.go ，尚未验证
4. 编写 cmd/api/rpc/comment.go 和 init.go ，尚未验证

需要配合完成的功能：在 cmd/comment/test 中的接口


# 20230201
基本框架搭建：
1. 编写idl/下的idl文件
2. 使用hertz命令行工具生成 cmd/api
3. 使用kitex命令行工具生成 cmd/comment 、 cmd/favorite 、 cmd/feed 、 cmd/message 、 cmd/publish 、cmd/relation 和 cmd/user
4. 使用kitex命令行工具生成kitex_gen
5. 创建 docker-compose.yml 和 pkg/configs/sql/init.sql 文件