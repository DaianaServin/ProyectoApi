package domain

type Turno struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	OdontologoID int    `json:"odontologo_id"`
	PacienteID int    `json:"paciente_id"`
}