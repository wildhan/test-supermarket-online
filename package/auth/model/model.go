package model

type UserAuth struct {
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

type ResponseLogin struct {
	Token string `json:"token" gorm:"column:token"`
}
