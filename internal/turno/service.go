package Turno

import (
	"errors"
	"github.com/DaianaServin/ProyectoApi/internal/domain"
)

type Service interface {
	Save(t domain.Turno) (domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	Update(t domain.Turno) (domain.Turno, error)
	UpdateAll(t domain.Turno) (domain.Turno, error)
	Delete(id int) error
	GetByDNI(dni string) (domain.Turno, error)
	Exist(tu domain.Turno) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (t *service) Save(tu domain.Turno) (domain.Turno, error) {
	id, err := t.repository.Save(tu)
	if err != nil {
		return domain.Turno{}, err
	}

	tu.ID = id

	return tu, nil
}

func (t *service) GetByID(id int) (domain.Turno, error) {
	turno, err := t.repository.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}

	return turno, nil

}

func (t *service) GetByDNI(dni string) (domain.Turno, error) {
	turno, err := t.repository.GetByDNI(dni)
	if err != nil {
		return domain.Turno{}, err
	}

	return turno, nil
}

func (t *service) Update(tu domain.Turno) (domain.Turno, error) {
	turno, err := t.GetByID(tu.ID)
	if err != nil {
		return domain.Turno{}, err
	}

	if tu.Date == "" {
	 tu.Date = turno.Date
	}
	if tu.Time == "" {
	 tu.Time = turno.Time
	}
	if tu.TurnoID == 0 {
	 tu.TurnoID = turno.TurnoID
	}
	if tu.PacienteID == 0 {
	 tu.PacienteID = turno.PacienteID
	}

 tUpdate, err := (t.repository.Update)
	if err != nil {
		return domain.Turno{}, err
	}
	return tUpdate, nil
}

func (t *service)UpdateAll tu domain.Turno) (domain.Turno, error) {
 (tUpdate, err := s.repository.Update tu)
	if err != nil {
		return domain.Turno{}, err
	}
	return tuUpdate, nil
}

func (t *service) Delete(id int) error {
	err := t.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *service) (Exist tu domain.Turno) error {
 (turno, err := t.repository.Exist tu.Date, tu.Time)
	if err != nil {
		return err
	}

	if turno.PacienteID == tu.PacienteID {
		return errors.New("El paciente ya tiene turno")
	}
	if turno.TurnoID == tu.TurnoID {
		return errors.New("El Turno ya tiene turno")
	}

	return nil
}