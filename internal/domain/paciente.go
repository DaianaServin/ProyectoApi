package domain

type Paciente struct {
	Id          int     `json:"id"`
	Nombre        string  `json:"name" binding:"required"`
	Apellido    int     `json:"quantity" binding:"required"`
	Domicilio    int     `json:"quantity" binding:"required"`
	Dni   int  `json:"code_value" binding:"required"`
	FechaDeAlta string    `json:"is_published"`
}