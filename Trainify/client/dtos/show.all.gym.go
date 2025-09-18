package dtos_user

type ShowAllGymsDTO struct {
	ID        uint    `json:"id"`
	Nombre    string  `json:"nombre"`
	Direccion string  `json:"direccion"`
	Telefono  string  `json:"telefono"`
	Foto      string  `json:"foto"`
	Latitud   float64 `json:"latitud"`
	Longitud  float64 `json:"longitud"`
}
