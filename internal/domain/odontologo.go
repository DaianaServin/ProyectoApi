package domain

type Odontologo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Enrollment string `json:"enrollment"`
}

func NewOdontologo(n string, sn string, e string) *Odontologo {
	return &Odontologo{
		ID:         0,
		Name:       n,
		Surname:    sn,
		Enrollment: e,
	}
}