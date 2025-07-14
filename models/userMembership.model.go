package models

import "time"

// UserMembership representa la afiliación de un usuario a un tipo de membresía
// Esta estructura rastrea el período de validez y estado de la membresía de un usuario
type UserMembership struct {
	// ID es el identificador único de la membresía del usuario
	ID uint `gorm:"primaryKey"`

	// UserID es el identificador del usuario que posee la membresía
	UserID uint

	// User es la relación con la entidad User
	User User `gorm:"foreignKey:UserID"`

	// MembershipTypeID es el identificador del tipo de membresía
	MembershipTypeID uint

	// MembershipType es la relación con la entidad MembershipType
	MembershipType MembershipType `gorm:"foreignKey:MembershipTypeID"`

	// StartDate es la fecha de inicio de la membresía
	StartDate time.Time `gorm:"type:timestamp"`

	// EndDate es la fecha de fin de la membresía
	EndDate time.Time `gorm:"type:timestamp"`

	// Active indica si la membresía está activa o no
	// Por defecto es true al crear una nueva membresía
	Active bool `gorm:"default:true"`
}
