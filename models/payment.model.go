package models

import "time"

// Payment representa un pago realizado por un usuario
// Esta estructura almacena todos los detalles de una transacción de pago
// incluyendo el monto, método de pago y estado de la transacción
type Payment struct {
	// ID es el identificador único del pago
	ID uint `gorm:"primaryKey"`

	// UserID es el identificador del usuario que realizó el pago
	UserID uint

	// User es la relación con la entidad User
	User User `gorm:"foreignKey:ID"`

	// GymID es el identificador del gimnasio donde se realizó el pago
	GymID uint

	// Gym es la relación con la entidad Gym
	Gym Gym `gorm:"foreignKey:ID"`

	// MembershipTypeID es el identificador del tipo de membresía pagada
	MembershipTypeID uint

	// MembershipType es la relación con la entidad MembershipType
	MembershipType MembershipType `gorm:"foreignKey:ID"`

	// Amount es el monto del pago
	Amount float64 `gorm:"type:decimal(10,2)"`

	// PaymentMethod es el método de pago utilizado
	// Puede ser 'tarjeta', 'efectivo' o 'transferencia'
	PaymentMethod string `gorm:"type:varchar(20)" validate:"required,oneof=tarjeta efectivo transferencia"`

	// Status es el estado actual del pago
	// Puede ser 'pendiente', 'completado' o 'rechazado'
	Status string `gorm:"type:varchar(20)" validate:"required,oneof=pendiente completado rechazado"`

	// Reference es el número de referencia del pago
	// Usado para identificar la transacción
	Reference string

	// CreatedAt es la fecha y hora de creación del registro del pago
	// Se establece automáticamente por GORM
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
