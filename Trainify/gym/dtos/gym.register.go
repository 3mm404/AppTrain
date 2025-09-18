package dtos_gym

import "time"

type RegisterGymDTO struct {
	Nombre    string    `json:"nombre" binding:"required"`
	Direccion string    `json:"direccion" binding:"required"`
	Telefono  string    `json:"telefono" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Foto      string    `json:"foto" binding:"omitempty,url"`
	Latitud   float64   `json:"latitud" binding:"required"`
	Longitud  float64   `json:"longitud" binding:"required"`
	Aprobado  bool      `json:"aprobado" binding:"omitempty"`
	Password  string    `json:"password" binding:"required,min=8,max=64"`
	CreadoEn  time.Time `json:"creado_en" binding:"omitempty"`
}

func (dto *RegisterGymDTO) SetDefaultValues() {
	// Si no se pasa "aprobado", que quede en false
	dto.Aprobado = false

	// Si no se pasa fecha de creación, que quede en la actual
	if dto.CreadoEn.IsZero() {
		dto.CreadoEn = time.Now()
	}
}
