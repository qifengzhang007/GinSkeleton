package test

import (
	"context"
	"ginskeleton/app/grpc/server/proto/stu_demo_pb"
	"ginskeleton/app/grpc/server/proto/user_demo_pb"
	"testing"

	"google.golang.org/grpc"
)

func TestGrpcClient1(t *testing.T) {
	//连接到gRPC服务端
	// err 永远返回 nil ，没有判断的必要
	conn, err := grpc.NewClient("127.0.0.1:20211", grpc.WithInsecure())
	defer conn.Close()

	stuServiceClient := stu_demo_pb.NewStudentServiceClient(conn)
	sendParams := &stu_demo_pb.StudentRequest{
		Name: "姓名关键词",
	}
	resp, err := stuServiceClient.GetStudentInfo(context.Background(), sendParams)
	if err != nil {
		t.Errorf("%s,%s", "GetStudentInfo - 接口调用失败", err.Error())
	} else {
		t.Logf("GetStudentInfo - 接口调用成功: %v", resp)
	}
}

func TestGrpcClient2(t *testing.T) {
	//连接到gRPC服务端
	conn, err := grpc.NewClient("127.0.0.1:20211", grpc.WithInsecure())
	defer conn.Close()

	userServiceClient := user_demo_pb.NewUserServiceClient(conn)
	sendParams := &user_demo_pb.UserRequest{
		Id: 1,
	}
	resp, err := userServiceClient.GetUserInfo(context.Background(), sendParams)

	if err != nil {
		t.Errorf("%s,%s", "GetUserInfo - 接口调用失败", err.Error())

	} else {
		t.Logf("GetUserInfo - 接口调用成功: %v", resp)
	}

}

func TestGrpcClient3(t *testing.T) {
	//连接到gRPC服务端
	conn, err := grpc.NewClient("127.0.0.1:20211", grpc.WithInsecure())
	defer conn.Close()

	userServiceClient := user_demo_pb.NewUserServiceClient(conn)
	sendParams := &user_demo_pb.UserIdRequest{
		Name: "姓名关键词",
	}
	resp, err := userServiceClient.GetItem(context.Background(), sendParams)

	if err != nil {
		t.Errorf("%s,%s", "GetItem - 接口调用失败", err.Error())
	} else {
		t.Logf("GetItem - 接口调用成功: %v", resp)
	}

}
