package dtos_user

type RegisterUserDTO struct {
	Nombre   string `json:"nombre" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Telefono string `json:"telefono" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}
