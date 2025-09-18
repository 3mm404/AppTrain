package dtos_gym

type ShowGymMoreMembershipDTO struct {
	ID           uint    `json:"id"`
	Gym          string  `json:"gym"` // ahora se verá como "gym": "Gym MaxPower"
	Descripcion  string  `json:"descripcion"`
	Precio       float64 `json:"precio"`
	DuracionDias int     `json:"duracion_dias"`
	Status       string  `json:"status"`
}
