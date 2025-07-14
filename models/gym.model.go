package models

import "time"

// Gym representa un gimnasio en el sistema
// Esta estructura almacena toda la información relevante sobre un gimnasio
// incluyendo sus datos de contacto, ubicación y estado de aprobación
type Gym struct {
	// ID es el identificador único del gimnasio
	ID uint `gorm:"primaryKey"`

	// Name es el nombre del gimnasio (máximo 100 caracteres)
	Name string `gorm:"type:varchar(100)"`

	// Address es la dirección física del gimnasio (máximo 255 caracteres)
	Address string `gorm:"type:varchar(255)"`

	// Phone es el número de teléfono del gimnasio (máximo 20 caracteres)
	Phone string `gorm:"type:varchar(20)"`

	// Email es la dirección de correo electrónico del gimnasio
	Email string `gorm:"type:varchar(100)"`

	// Photo es la URL o ruta de la imagen del gimnasio
	Photo string `gorm:"type:varchar(255)"`

	// Approved indica si el gimnasio ha sido aprobado por el administrador
	// Por defecto es false hasta que sea revisado
	Approved bool `gorm:"default:false"`

	// Latitude es la coordenada de latitud de la ubicación del gimnasio
	Latitude float64 `gorm:"type:decimal(10,6)"`

	// Longitude es la coordenada de longitud de la ubicación del gimnasio
	Longitude float64 `gorm:"type:decimal(10,6)"`

	// Password es la contraseña del gimnasio para acceso al sistema
	Password string `gorm:"type:varchar(255)"`

	// CreatedAt es la fecha y hora de creación del registro del gimnasio
	// Se establece automáticamente por GORM
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
