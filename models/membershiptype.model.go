package models

import "time"

// MembershipType representa un tipo de membresía disponible en un gimnasio
// Esta estructura define los diferentes planes de membresía que un gimnasio puede ofrecer
type MembershipType struct {
	// ID es el identificador único del tipo de membresía
	ID uint `gorm:"primaryKey"`

	// GymID es el identificador del gimnasio que ofrece este tipo de membresía
	GymID uint

	// Gym es la relación con la entidad Gym
	Gym Gym `gorm:"foreignKey:ID"`

	// Name es el nombre del tipo de membresía (máximo 100 caracteres)
	Name string `gorm:"type:varchar(100)"`

	// Description es la descripción detallada del tipo de membresía
	Description string `gorm:"type:text"`

	// Price es el precio del tipo de membresía
	Price float64 `gorm:"type:decimal(10,2)"`

	// DurationDays es la duración en días de este tipo de membresía
	DurationDays int `gorm:"type:int"`

	// Status indica si el tipo de membresía está activo o no
	// Por defecto es 'activo'
	Status string `gorm:"type:varchar(20);default:'activo'"`

	// CreatedAt es la fecha y hora de creación del registro del tipo de membresía
	// Se establece automáticamente por GORM
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
