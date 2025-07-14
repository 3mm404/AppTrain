package dtos

import "time"

type RegisterUserDTO struct {
	Nombre             string    `json:"nombre" binding:"required"`
	Email              string    `json:"email" binding:"required,email"`
	Telefono           string    `json:"telefono" binding:"required_if=Email,0"`
	FechaNacimiento    time.Time `json:"fecha_nacimiento" binding:"omitempty"`
	Password           string    `json:"password" binding:"required,min=8,max=16"`
	Foto               string    `json:"foto" binding:"omitempty,url"`
	TipoUsuario        string    `json:"tipo_usuario" binding:"omitempty,oneof=cliente empleado"`
	CodigoVerificacion string    `json:"codigo_verificacion" binding:"omitempty,min=6,max=6"`
	CreadoEn           time.Time `json:"creado_en" binding:"omitempty"`
}

func (dto *RegisterUserDTO) SetDefaultValues() {
	if dto.TipoUsuario == "" {
		dto.TipoUsuario = "cliente" // Default to "cliente"
	}
}
