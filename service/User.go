package service

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Register 用户注册
func (service *UserRegisterService) Register() {
	//user := model.User{
	//	Nickname: service.Nickname,
	//	UserName: service.UserName,
	//	Status:   model.Active,
	//}
	//
	//
	//// 加密密码
	//
	//
	//// 创建用户
	//if err := model.DB.Create(&user).Error; err != nil {
	//	return user, &serializer.Response{
	//		Status: 40002,
	//		Msg:    "注册失败",
	//	}
	//}
	//
}
