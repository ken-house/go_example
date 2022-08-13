package model

// LoginForm 登录表单数据
type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,max=30"`
	Password string `form:"password" json:"password" binding:"required,validatePassword"`
}
