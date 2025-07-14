package models

import "time"

// Attendance representa el registro de entrada y salida de un usuario en un gimnasio
// Esta estructura se utiliza para rastrear la asistencia de los usuarios
// y generar reportes de uso del gimnasio
type Attendance struct {
	// ID es el identificador único del registro de asistencia
	ID uint `gorm:"primaryKey"`

	// UserID es el identificador del usuario que registró la asistencia
	// Se relaciona con la tabla users a través de una clave foránea
	UserID uint `gorm:"foreignKey:ID"`

	// GymID es el identificador del gimnasio donde se registró la asistencia
	// Se relaciona con la tabla gyms a través de una clave foránea
	GymID uint `gorm:"foreignKey:ID"`

	// CheckIn es la fecha y hora de entrada del usuario al gimnasio
	CheckIn time.Time `gorm:"type:timestamp"`

	// CheckOut es la fecha y hora de salida del usuario del gimnasio
	// Este campo puede ser nulo si el usuario aún no ha salido
	CheckOut time.Time `gorm:"type:timestamp;null"`
}
