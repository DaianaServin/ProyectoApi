package Odontologo

import (
	"github.com/DaianaServin/ProyectoApi/internal/domain"
	"errors"
)

type Service interface {
	Save(d domain.Odontologo) (domain.Odontologo, error)
	GetByID(id int) (domain.Odontologo, error)
	Update(d domain.Odontologo) (domain.Odontologo, error)
	UpdateAll(d domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
	Exists(enrollment string) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(d domain.Odontologo) (domain.Odontologo, error) {
	newodontologo := domain.Newodontologo(d.Name, d.Surname, d.Enrollment)

	id, err := s.repository.Save(*newodontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}

	newodontologo.ID = id

	return *newodontologo, nil
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil

}

func (s *service) Update(d domain.Odontologo) (domain.Odontologo, error) {
	odontologo, err := s.GetByID(d.ID)
	if err != nil {
		return domain.Odontologo{}, err
	}

	if d.Name == "" {
		d.Name = odontologo.Name
	}
	if d.Surname == "" {
		d.Surname = odontologo.Surname
	}
	if d.Enrollment == "" {
		d.Enrollment = odontologo.Enrollment
	}

	odontologoToUpdate := domain.NewOdontologo(d.Name, d.Surname, d.Enrollment)
	odontologoToUpdate.ID = d.ID

	oUpdate, err := s.repository.Update(*odontologoToUpdate)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return oUpdate, nil
}

func (s *service) UpdateAll(d domain.Odontologo) (domain.Odontologo, error) {
	dUpdate, err := s.repository.Update(d)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return dUpdate, nil
}

func (s *service) Delete(id int) error {
	hasTurno := s.repository.HasShifts(id)
	if hasTurno {
		return errors.New("El odontologo tiene un turno asignado")
	}

	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Exists(enrollment string) bool {
	exist := s.repository.Exists(enrollment)
	return exist
}