package main

import (
	"ginskeleton/app/global/variable"
	grpc_interceptor "ginskeleton/app/grpc/server/interceptor"
	_ "ginskeleton/bootstrap"
	"ginskeleton/routers"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//1.指定执行程序监听的端口
	lis, err := net.Listen("tcp", variable.ConfigYml.GetString("GrpcServer.Port"))
	if err != nil {
		log.Fatalf("Tcp 监听失败: %v", err)
	}

	//2.初始化 gPRC 服务，并注册服务
	grpcServ := grpc.NewServer(grpc.UnaryInterceptor(grpc_interceptor.GrpcRequestLog()))
	//  业务服务注册
	routers.InitGrpcService(grpcServ)
	variable.ZapLog.Info("开始启动 grpc 服务, 监听端口: " + variable.ConfigYml.GetString("GrpcServer.Port"))
	//3.启动服务
	if err = grpcServ.Serve(lis); err != nil {
		log.Fatalf("grpc 服务启动失败,错误: %v", err)
	}

}
