package Users

type Base struct {
	Name string `form:"name" json:"name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}
