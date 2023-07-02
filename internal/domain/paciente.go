package domain

type Paciente struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Address       string `json:"address"`
	DNI           string `json:"DNI"`
	TurnoFecha string `json:"discharge_date"`
}

func NewPaciente(name string, surname string, address string, dni string, fechaTurno string) *Paciente {
	return &Paciente{
		ID:            0,
		Name:          name,
		Surname:       surname,
		Address:       address,
		DNI:           dni,
		TurnoFecha: fechaTurno,
	}
}