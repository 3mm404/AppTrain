package models

import "time"

type Payment struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	User             User `gorm:"foreignKey:UserID"`
	GymID            uint
	Gym              Gym `gorm:"foreignKey:GymID"`
	MembershipTypeID uint
	MembershipType   MembershipType `gorm:"foreignKey:MembershipTypeID"`
	Monto            float64        `gorm:"type:decimal(10,2)"`
	MetodoPago       string         `gorm:"type:varchar(20)" validate:"required,oneof=tarjeta efectivo transferencia"`
	Estado           string         `gorm:"type:varchar(20)" validate:"required,oneof=pendiente completado rechazado"`
	Referencia       string
	CreadoEn         time.Time `gorm:"autoCreateTime"`
}
