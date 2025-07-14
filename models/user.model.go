package models

import "time"

// User representa un usuario del sistema
// Esta estructura almacena toda la información relevante de un usuario
// incluyendo sus datos personales y su rol en el sistema
type User struct {
	// ID es el identificador único del usuario
	ID uint `gorm:"primaryKey"`

	// Name es el nombre completo del usuario (máximo 100 caracteres)
	Name string `gorm:"type:varchar(100)"`

	// Email es la dirección de correo electrónico del usuario
	// Debe ser única en el sistema
	Email string `gorm:"type:varchar(100);unique"`

	// Phone es el número de teléfono del usuario
	// Debe ser único en el sistema y se puede usar para login
	Phone string `gorm:"type:varchar(15);unique"`

	// Password es la contraseña del usuario (hash)
	Password string `gorm:"type:varchar(255)"`

	// Photo es la URL o ruta de la imagen de perfil del usuario
	Photo string `gorm:"type:varchar(255)"`

	// BirthDate es la fecha de nacimiento del usuario
	BirthDate time.Time `gorm:"type:date"`

	// UserType indica el tipo de usuario
	// Puede ser 'cliente' o 'empleado'
	// Por defecto es 'cliente'
	UserType string `gorm:"type:varchar(20);default:'cliente'"`

	// GymID es el identificador del gimnasio donde trabaja el usuario
	// Solo se usa para empleados
	GymID *uint `gorm:"index"`

	// Gym es la relación con la entidad Gym
	// Solo se usa para empleados
	Gym *Gym `gorm:"foreignKey:ID"`

	// VerificationCode es el código de verificación para SMS
	VerificationCode string `gorm:"type:varchar(6)"`

	// CreatedAt es la fecha y hora de creación del registro del usuario
	// Se establece automáticamente por GORM
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
