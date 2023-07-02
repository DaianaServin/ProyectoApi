package Turno

import (
	"database/sql"
	"fmt"
	"github.com/DaianaServin/ProyectoApi/internal/domain"
)

type Repository interface {
	Save(t domain.Turno) (int, error)
	GetByID(id int) (domain.Turno, error)
	Update(domain.Turno) (domain.Turno, error)
	Delete(id int) error
	GetByDNI(dni string) (domain.Turno, error)
	Exist(data string, time string) (domain.Turno, error)
}

const (
	SAVE_TURNO         = "INSERT INTO turnos(data, time, odontologo_id, paciente_id) VALUES (?, ?, ?, ?);"
	GET_TURNO_BY_ID    = "SELECT * FROM turnos WHERE id = ?;"
	GET_TURNO_BY_DNI   = "SELECT t.* FROM turnos t INNER JOIN pacientes p ON p.id = t.paciente_id where p.dni = ? GROUP BY p.dni;"
	UPDATE_TURNO       = "UPDATE turnos SET data = ?, time = ?, odontologo_id = ?, paciente_id = ? WHERE id = ?;"
	DELETE_TURNO_BY_ID = "DELETE FROM turnos WHERE id = ?;"
	EXIST              = "SELECT * FROM turnos WHERE data = ? AND time = ?;"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(t domain.Turno) (int, error) {
	stmt, err := r.db.Prepare(SAVE_TURNO)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&t.Date, &t.Time, &t.odontologoID, &t.pacienteID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(id int) (domain.Turno, error) {
	row := r.db.QueryRow(GET_TURNO_BY_ID, id)
	t := domain.Turno{}
	err := row.Scan(&t.ID, &t.Date, &t.Time, &t.odontologoID, &t.pacienteID)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (r *repository) GetByDNI(dni string) (domain.Turno, error) {
	t := domain.Turno{}

	row, err := r.db.Query(GET_TURNO_BY_DNI, dni)
	if err != nil {
		return t, err
	}

	for row.Next() {
		err := row.Scan(&t.ID, &t.Date, &t.Time, &t.odontologoID, &t.pacienteID)
		if err != nil {
			return t, err
		}

	}

	return t, nil
}

func (r *repository) Update(t domain.Turno) (domain.Turno, error) {
	stmt, err := r.db.Prepare(UPDATE_TURNO)
	if err != nil {
		return domain.Turno{}, err
	}

	res, err := stmt.Exec(&t.Date, &t.Time, &t.odontologoID, &t.pacienteID, &t.ID)
	if err != nil {
		return domain.Turno{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Turno{}, err
	}

	return t, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DELETE_TURNO_BY_ID)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowAffect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffect < 1 {
		return fmt.Errorf("NotFound")
	}

	return nil
}

func (r *repository) Exist(data string, time string) (domain.Turno, error) {
	row, _ := r.db.Query(EXIST, data, time)
	t := domain.Turno{}
	for row.Next() {
		err := row.Scan(&t.ID, &t.Date, &t.Time, &t.odontologoID, &t.pacienteID)
		if err != nil {
			return t, err
		}
	}

	return t, nil
}