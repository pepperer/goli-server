package api

// 用于获取请求的参数 模型

type RegisterParams struct {
	Username string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Token    string `form:"token" json:"token" binding:"required"`
}
