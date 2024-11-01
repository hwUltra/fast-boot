#!/bin/bash
#进目录 创建Dockerfile
#goctl docker -go sys.go

#构建服务
docker build -t rpc-sys:v1 -f app/rpc/sys/Dockerfile .
docker build -t api-admin:v1 -f app/api/admin/Dockerfile .

#启动服务
docker run -itd --net=host --name=rpc-sys rpc-sys:v1
docker run -itd --net=host --name=api-api api-admin:v1

#停止服务
docker stop rpc-sys
docker stop api-admin

#删除容器
docker rm rpc-sys
docker rm api-admin

#删除镜像
docker rmi rpc-sys:v1
docker rmi api-admin:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')
