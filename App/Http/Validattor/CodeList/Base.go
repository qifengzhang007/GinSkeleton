package CodeList

type CodelistBase struct {
	Code string `form:"code" json:"code"  bind:"required"`
	Name string `form:"name" json:"name"  bind:"required"`
}
