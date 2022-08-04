package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"id" binding:"required"`
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
}
