package curd

import (
	"goskeleton/app/model"
	"goskeleton/app/utils/md5_encrypt"
)

func CreateUserCurdFactory() *UsersCurd {

	return &UsersCurd{}
}

type UsersCurd struct {
}

func (u *UsersCurd) Register(data *model.UsersModel) bool {
	data.Pass = md5_encrypt.Base64Md5(data.Pass) // 预先处理密码加密，然后存储在数据库
	err := model.CreateUserFactory("").Register(data)
	if err != nil {
		return false
	}
	return true
}

func (u *UsersCurd) Store(name string, pass string, realName string, phone string, remark string) bool {

	pass = md5_encrypt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return model.CreateUserFactory("").Store(name, pass, realName, phone, remark)
}

func (u *UsersCurd) Update(id float64, name string, pass string, realName string, phone string, remark string, clientIp string) bool {
	//预先处理密码加密等操作，然后进行更新
	pass = md5_encrypt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return model.CreateUserFactory("").Update(id, name, pass, realName, phone, remark, clientIp)
}
