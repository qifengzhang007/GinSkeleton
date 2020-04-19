package Curd

import "GinSkeleton/App/Model"

func CreateUserCurdFactory() *UsersCurd {

	return &UsersCurd{}
}

type UsersCurd struct {
}

func (u *UsersCurd) Store(name string, pass string, real_name string, phone string, remark string) bool {

	// 预先处理密码加密，然后存储在数据库
	return Model.CreateUserFactory().Store(name, pass, real_name, phone, remark)
}

func (u *UsersCurd) Update(id float64, name string, pass string, real_name string, phone string, remark string) bool {
	//预先处理密码加密等操作，然后进行更新
	// 更新之前有需要预先处理的业务，走一层service，否则直接调用mode中的update即可
	return Model.CreateUserFactory().Update(id, name, pass, real_name, phone, remark)
}
