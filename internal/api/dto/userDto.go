package dto

// 接收参数结构体

// UserLogin 用户登录
type UserLogin struct {
	UserName string `json:"username" binding:"max=20"`
	PassWord string `json:"password" binding:"required"`
}

// ThirdCallback 第三方登录回调参数
type ThirdCallback struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

// EmailCode 邮件验证码
type EmailCode struct {
	Email string `json:"email" binding:"required,email"`
}

// UserRegister 注册
type UserRegister struct {
	UserLogin
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

// RefreshRequest 刷新token请求参数
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserInfo 获取用户信息参数
type UserInfo struct {
	UserId int `json:"userId,omitempty" binding:"required"`
}

// UpdatePass 修改用户密码参数
type UpdatePass struct {
	Code     string `json:"code,omitempty" binding:"required,len=6"`
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
}

// UpdateUser 更新用户信息参数
type UpdateUser struct {
	UserName string `json:"user_name,omitempty" binding:"max=20"`
	Email    string `json:"email,omitempty" binding:"required,email"`
	Role     string `json:"role,omitempty" binding:"required"`
	Avatar   string `json:"avatar,omitempty" binding:"required"`
}
