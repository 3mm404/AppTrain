package dtos_user

import "time"

type InfMembershipClient struct {
	GymID        uint      `json:"gym_id"` // ID del gimnasio al que pertenece la membresía
	GymName      string    `json:"gym_name"` // Nombre del gimnasio
	Status       string    `json:"status"` // "activo" o "inactivo"
	StartDate    time.Time `json:"start_date"` // Fecha de inicio de la membresía
	EndDate      time.Time `json:"end_date"`   // Fecha de expiración de la membresía
}
