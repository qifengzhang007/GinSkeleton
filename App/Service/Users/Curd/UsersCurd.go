package Curd

import (
	"GinSkeleton/App/Model"
	"GinSkeleton/App/Utils/MD5Cryt"
)

func CreateUserCurdFactory() *UsersCurd {

	return &UsersCurd{}
}

type UsersCurd struct {
}

func (u *UsersCurd) Register(name string, pass string, user_ip string) bool {
	pass = MD5Cryt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return Model.CreateUserFactory("").Register(name, pass, user_ip)
}

func (u *UsersCurd) Store(name string, pass string, real_name string, phone string, remark string) bool {

	pass = MD5Cryt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return Model.CreateUserFactory("").Store(name, pass, real_name, phone, remark)
}

func (u *UsersCurd) Update(id float64, name string, pass string, real_name string, phone string, remark string, client_ip string) bool {
	//预先处理密码加密等操作，然后进行更新
	pass = MD5Cryt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return Model.CreateUserFactory("").Update(id, name, pass, real_name, phone, remark, client_ip)
}
