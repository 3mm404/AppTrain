package models

import "time"

// GymEmployee representa un empleado de un gimnasio
// Esta estructura almacena la relación entre un usuario y su rol en un gimnasio
// y su estado actual (activo/inactivo)
type GymEmployee struct {
	// ID es el identificador único del empleado del gimnasio
	ID uint `gorm:"primaryKey"`

	// GymID es el identificador del gimnasio donde trabaja el empleado
	GymID uint `gorm:"foreignKey:ID"`

	// UserID es el identificador del usuario que es empleado del gimnasio
	UserID uint `gorm:"foreignKey:ID"`

	// Role es el rol del empleado en el gimnasio
	// Puede ser 'admin' o 'empleado'
	// Por defecto es 'empleado'
	Role string `gorm:"type:varchar(20);default:'empleado'" validate:"required,oneof=admin empleado"`

	// Status indica si el empleado está activo o inactivo
	// Puede ser 'activo' o 'inactivo'
	// Por defecto es 'activo'
	Status string `gorm:"type:varchar(20);default:'activo'" validate:"required,oneof=activo inactivo"`

	// CreatedAt es la fecha y hora de creación del registro del empleado
	// Se establece automáticamente por GORM
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
