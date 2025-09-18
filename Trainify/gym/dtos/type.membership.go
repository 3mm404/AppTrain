package dtos_gym

// MembershipTypeDTO define los campos necesarios para crear un tipo de membresía
type MembershipTypeDTO struct {
	ID           uint    `json:"id"`
	Nombre       string  `json:"nombre" validate:"required"`
	Descripcion  string  `json:"descripcion" validate:"required"`
	Precio       float64 `json:"precio" validate:"required,gt=0"`
	DuracionDias int     `json:"duracion_dias" validate:"required,gt=0"`
	Status       string  `json:"status" validate:"required,oneof=activo inactivo"`
}
