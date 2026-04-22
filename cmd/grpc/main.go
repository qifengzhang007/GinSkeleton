package main

import (
	"ginskeleton/app/global/variable"
	"ginskeleton/app/grpc/server/proto/stu_demo_pb"
	"ginskeleton/app/grpc/server/proto/user_demo_pb"
	"ginskeleton/app/grpc/server/service_implement"
	_ "ginskeleton/bootstrap"
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

	//2.建立 gPRC 服务器，并注册服务
	grpcServ := grpc.NewServer()
	// PB 文件调用注册函数，将grpc与业务service进行绑定、注册
	user_demo_pb.RegisterUserServiceServer(grpcServ, &service_implement.UserService{})
	stu_demo_pb.RegisterStudentServiceServer(grpcServ, &service_implement.StuService{})

	variable.ZapLog.Info("Grpc 服务启动, 监听端口: " + variable.ConfigYml.GetString("GrpcServer.Port"))
	//3.启动服务
	if err = grpcServ.Serve(lis); err != nil {
		log.Fatalf("Grpc 服务启动失败,错误: %v", err)
	}

}
