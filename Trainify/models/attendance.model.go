package models

import "time"

type Attendance struct {
	ID      uint      `gorm:"primaryKey"`
	UserID  uint      `gorm:"foreignKey"`
	GymID   uint      `gorm:"foreignKey"`
	Entrada time.Time `gorm:"type:timestamp"`
	Salida  time.Time `gorm:"type:timestamp"`
}
