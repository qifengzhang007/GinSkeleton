package Users

type Base struct {
	Name string `form:"username" json:"username"  bind:"required"`
}
