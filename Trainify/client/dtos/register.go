package dtos_user

type ReUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Telefono string `json:"telefono" binding:"required_if=Email,0"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}
