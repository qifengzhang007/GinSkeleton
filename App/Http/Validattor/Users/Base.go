package Users

type Base struct {
	Name string `form:"name" json:"name"  bind:"required"`
}
