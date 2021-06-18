package users

type BaseField struct {
	UserName string `form:"user_name" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Pass     string `form:"pass" json:"pass" binding:"required,min=6,max=20"`     //  密码为 必填，长度>=6
}

type Id struct {
	Id float64 `form:"id"  json:"id" binding:"required,min=1"`
}
