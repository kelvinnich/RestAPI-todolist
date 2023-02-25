package dto

type RegisterUsersDTO struct {
	ID string `json:"id"`
	Username string `json:"username" form:"username"binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Address string `json:"address" form:"address"`
}

type LoginUsersDTO struct {
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUsersDTO struct {
	ID string `json:"id"`
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address string `json:"address"`
}