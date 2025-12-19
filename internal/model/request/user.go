package request

// 接收参数结构体

// UserRegister 注册
type UserRegister struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required"`
}

// EmailCode 邮件验证码
type EmailCode struct {
	Email string `json:"email" binding:"required,email"`
}

// UserLogin 用户登录
type UserLogin struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

// RefreshRequest 刷新token请求参数
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UpdateUser 更新用户信息参数
type UpdateUser struct {
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
