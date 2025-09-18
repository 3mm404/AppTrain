package models

import "time"

type UserMembership struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	User             User `gorm:"foreignKey:UserID"` // Asociación con User
	MembershipTypeID uint
	MembershipType   MembershipType `gorm:"foreignKey:MembershipTypeID"` // Asociación con MembershipType
	FechaInicio      time.Time      `gorm:"type:timestamp"`
	FechaFin         time.Time      `gorm:"type:timestamp"`
	Activo           bool           `gorm:"default:true"`
}
