package routers

import (
	"ginskeleton/app/grpc/server/proto/stu_demo_pb"
	"ginskeleton/app/grpc/server/proto/user_demo_pb"
	"ginskeleton/app/grpc/server/service_implement"

	"google.golang.org/grpc"
)

func InitGrpcService(grpcServ *grpc.Server) {
	// PB 文件调用注册函数，将grpc与业务service进行绑定、注册
	// 1.注册 UserService示例服务
	user_demo_pb.RegisterUserServiceServer(grpcServ, &service_implement.UserService{})
	// 2.注册 StudentService示例服务
	stu_demo_pb.RegisterStudentServiceServer(grpcServ, &service_implement.StuService{})
}
