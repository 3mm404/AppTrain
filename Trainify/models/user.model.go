package models

import "time"

type User struct {
	ID                 uint      `gorm:"primaryKey"`
	Nombre             string    `gorm:"type:varchar(100)"`
	Email              string    `gorm:"type:varchar(100);unique"`
	Telefono           string    `gorm:"type:varchar(15);unique"` // Para login con número
	Password           string    `gorm:"type:varchar(255)"`
	Foto               string    `gorm:"type:varchar(255)"`
	FechaNacimiento    time.Time `gorm:"type:date"`
	TipoUsuario        string    `gorm:"type:varchar(20);default:'cliente'"` // 'cliente' o 'empleado'
	GymID              *uint     `gorm:"index"`                              // Solo si es empleado
	Gym                *Gym      `gorm:"foreignKey:GymID"`                   // Relación con gimnasio (solo para empleados)
	CodigoVerificacion string    `gorm:"type:varchar(6)"`                    // Para autenticación SMS
	CreadoEn           time.Time `gorm:"autoCreateTime"`
}
