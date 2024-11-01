1.安装
框架安装 go get -u github.com/zeromicro/go-zero
框架代码生成工具安装 go get -u github.com/zeromicro/go-zero/tools/goctl

2.创建api
进到/app/api/admin/desc/目录执行
goctl api -o admin.api
goctl api go -api admin.api -dir ../

3.创建rpc
进到rpc/sys/目录操作
goctl rpc template -o sys.proto
goctl rpc protoc sys.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

4.创建model
进到rpc/目录操作
goctl model mysql ddl -c -src all.sql -dir .
goctl model mysql datasource -url="root:123456@tcp(192.168.3.88:3306)/gozero" -table="sys*" -dir ./model/sysmodel