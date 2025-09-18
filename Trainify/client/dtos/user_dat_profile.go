package dtos_user

import "time"

type ShowInformationUserProfile struct {
	Nombre          string    `json:"nombre" binding:"required"`
	Email           string    `json:"email" binding:"required,email"`
	Telefono        string    `json:"telefono" binding:"required_if=Email,0"`
	FechaNacimiento time.Time `json:"fecha_nacimiento" binding:"omitempty"`
	Password        string    `json:"password" binding:"required,min=8,max=16"`
	Foto            string    `json:"foto" binding:"omitempty,url"`
}
