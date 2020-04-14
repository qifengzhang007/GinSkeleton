package Users

type UsersRegister struct {
	Base
	Phone string `form:"phone" json:"phone"  bind:"required"`
	Pass  string `form:"pass" json:"pass bind:"required""`
}
