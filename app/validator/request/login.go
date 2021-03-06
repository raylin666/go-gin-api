package request

type LoginValidate struct {
	Username string `validate:"required" label:"用户名"`
	Password string `validate:"required" label:"用户密码"`
}