package service_implement

import (
	"context"
	"ginskeleton/app/grpc/server/proto/stu_demo_pb"
)

type StuService struct {
	// 定义一个测试函方法
}

func (u *StuService) GetStudentInfo(ctx context.Context, req *stu_demo_pb.StudentRequest) (resp *stu_demo_pb.StudentResponse, err error) {

	// 这里直接返回固定的数据，快速演示 grpc-service 功能即可
	resp = &stu_demo_pb.StudentResponse{
		Code: 200,
		Msg:  "success",
		Data: []*stu_demo_pb.Data{
			{
				Id:     1,
				Name:   req.Name + " - 测试姓名001",
				Age:    18,
				School: "xxx  - 小学001",
			},
			{
				Id:     2,
				Name:   req.Name + " - 测试姓名002",
				Age:    19,
				School: "xxx  - 小学002",
			},
		},
	}
	return resp, nil
}
