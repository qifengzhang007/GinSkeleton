package service_implement

import (
	"context"

	"ginskeleton/app/grpc/server/proto/user_demo_pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/anypb"
)

type UserService struct {
	// 定义一个测试函方法
}

func (u *UserService) GetUserInfo(ctx context.Context, req *user_demo_pb.UserRequest) (resp *user_demo_pb.UserResponse, err error) {
	if err = req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "GetUserInfo - 接口请求的参数校验不通过: %v", err)
	}
	resp = &user_demo_pb.UserResponse{
		Code: 200,
		Msg:  "success",
		Data: []*user_demo_pb.Data2{
			&user_demo_pb.Data2{
				OneItem: map[string]*anypb.Any{
					"user_name": {
						TypeUrl: "type.googleapis.com/google.protobuf.StringValue", // 指定消息类型的 URL
						Value:   []byte("zhangsan001"),                             // 存储编码后的消息字节流
					},
					"age": {
						TypeUrl: "type.googleapis.com/google.protobuf.Int32Value", // 指定消息类型的 URL
						Value:   []byte("18"),                                     // 存储编码后的消息字节流
					},
					"real_name": {
						TypeUrl: "type.googleapis.com/google.protobuf.StringValue", // 指定消息类型的 URL
						Value:   []byte("张三2026"),                                  // 存储编码后的消息字节流
					},
				},
			},
			&user_demo_pb.Data2{
				OneItem: map[string]*anypb.Any{
					"user_name": {
						TypeUrl: "type.googleapis.com/google.protobuf.StringValue", // 指定消息类型的 URL
						Value:   []byte("lisi002"),                                 // 存储编码后的消息字节流
					},
					"age": {
						TypeUrl: "type.googleapis.com/google.protobuf.Int32Value", // 指定消息类型的 URL
						Value:   []byte("19"),                                     // 存储编码后的消息字节流
					},
					"real_name": {
						TypeUrl: "type.googleapis.com/google.protobuf.StringValue", // 指定消息类型的 URL
						Value:   []byte("李四2026"),                                  // 存储编码后的消息字节流
					},
				},
			},
		},
	}
	return resp, nil
}

func (u *UserService) GetItem(ctx context.Context, req *user_demo_pb.UserIdRequest) (resp *user_demo_pb.UserResponse, err error) {
	if err = req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "GetItem - 接口请求的参数校验不通过: %v", err)
	}

	resp = &user_demo_pb.UserResponse{
		Code: 200,
		Msg:  "success",
		Data: []*user_demo_pb.Data2{
			&user_demo_pb.Data2{
				OneItem: map[string]*anypb.Any{
					"user_name": {
						TypeUrl: "type.googleapis.com/google.protobuf.StringValue", // 指定消息类型的 URL
						Value:   []byte("zhangsan"),                                // 存储编码后的消息字节流
					},
					"age": {
						TypeUrl: "type.googleapis.com/google.protobuf.Int32Value", // 指定消息类型的 URL
						Value:   []byte("18"),                                     // 存储编码后的消息字节流
					},
					"real_name": {
						TypeUrl: "type.googleapis.com/google.protobuf.StringValue", // 指定消息类型的 URL
						Value:   []byte("张三2026"),                                  // 存储编码后的消息字节流
					},
				},
			},
		},
	}
	return resp, nil
}
