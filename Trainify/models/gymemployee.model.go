package models

import "time"

type GymEmployee struct {
	ID       uint      `gorm:"primaryKey"`
	GymID    uint      `gorm:"foreignKey"`
	UserID   uint      `gorm:"foreignKey"`
	Rol      string    `gorm:"type:varchar(20);default:'empleado'" validate:"required,oneof=admin empleado"`
	Estado   string    `gorm:"type:varchar(20);default:'activo'" validate:"required,oneof=activo inactivo"`
	CreadoEn time.Time `gorm:"autoCreateTime"`
}
