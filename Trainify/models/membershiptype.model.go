package models

import "time"

type MembershipType struct {
	ID           uint      `gorm:"primaryKey"`
	GymID        uint      // Relación con Gym
	Gym          Gym       `gorm:"foreignKey:GymID"` // Asociación con la tabla Gym
	Nombre       string    `gorm:"type:varchar(100)"`
	Descripcion  string    `gorm:"type:text"`
	Precio       float64   `gorm:"type:decimal(10,2)"`
	DuracionDias int       `gorm:"type:int"`
	Status       string    `gorm:"type:varchar(20);default:'activo'"` // Agregamos el campo Status con un valor por defecto
	CreadoEn     time.Time `gorm:"autoCreateTime"`
}
