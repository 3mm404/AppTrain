package dtos_user

type LoginUserDTO struct {
	Identificador string `json:"identificador"` // Puede ser email o teléfono
	Password      string `json:"password"`
}
