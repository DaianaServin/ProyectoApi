package domain

type Odontologo struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Apellido    string     `json:"quantity" binding:"required"`
	Matricula    int     `json:"quantity" binding:"required"`
}