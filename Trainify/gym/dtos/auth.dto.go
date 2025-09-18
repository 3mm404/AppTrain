package dtos_gym

type LoginGymDTO struct {
	Identificador string `json:"identificador"` // Puede ser email o teléfono
	Password      string `json:"password"`
}
