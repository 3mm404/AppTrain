package models

import "time"

type Gym struct {
	ID        uint      `gorm:"primaryKey"`
	Nombre    string    `gorm:"type:varchar(100)"`
	Direccion string    `gorm:"type:varchar(255)"`
	Telefono  string    `gorm:"type:varchar(20)"`
	Email     string    `gorm:"type:varchar(100)"`
	Foto      string    `gorm:"type:varchar(255)"`
	Aprobado  bool      `gorm:"default:false"`
	Latitud   float64   `gorm:"type:decimal(10,6)"`
	Longitud  float64   `gorm:"type:decimal(10,6)"`
	Password  string    `gorm:"type:varchar(255)"`
	CreadoEn  time.Time `gorm:"autoCreateTime"`
}
