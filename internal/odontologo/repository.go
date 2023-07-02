package odontologo

import (
	"github.com/DaianaServin/ProyectoApi/internal/domain"
	"database/sql"
	"fmt"
)

type Repository interface {
	Save(o domain.Odontologo) (int, error)
	GetByID(id int) (domain.Odontologo, error)
	Update(o domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
	Exists(enrollment string) bool
	HasTurno(id int) bool
}

const (
	SAVE_ODONTOLOGO         = "INSERT INTO odontologos(name,surname,enrollment) VALUES (?,?,?);"
	GET_ODONTOLOGO_BY_ID    = "SELECT * FROM odontologos WHERE id = ?;"
	UPDATE_ODONTOLOGO       = "UPDATE odontologos SET name = ?, surname = ?, enrollment = ? WHERE id = ?;"
	DELETE_ODONTOLOGO_BY_ID = "DELETE FROM odontologos WHERE id = ?;"
	EXIST_ODONTOLOGO        = "SELECT enrollment FROM odontologos WHERE enrollment = ?"
	HAS_TURNO            = "SELECT o.* FROM odontologos o INNER JOIN turnos t ON o.id = t.odontologo_id WHERE odontologo_id = ?;"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,	}
}

func (r *repository) Save(o domain.Odontologo) (int, error) {
	stmt, err := r.db.Prepare(SAVE_ODONTOLOGO)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&o.Name, &o.Surname, &o.Enrollment)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(id int) (domain.Odontologo, error) {
	row := r.db.QueryRow(GET_ODONTOLOGO_BY_ID, id)
	O := domain.Odontologo{}
	err := row.Scan(&o.ID, &o.Name, &o.Surname, &o.Enrollment)
	if err != nil {
		return o, err
	}

	return o, nil
}

func (r *repository) Update(o domain.Odontologo) (domain.Odontologo, error) {
	stmt, err := r.db.Prepare(UPDATE_ODONTOLOGO)
	if err != nil {
		return domain.Odontologo{}, err
	}

	res, err := stmt.Exec(&o.Name, &o.Surname, &o.Enrollment, &o.ID)
	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	return o, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DELETE_ODONTOLOGO_BY_ID)
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

func (r *repository) Exists(enrollment string) bool {
	row := r.db.QueryRow(EXIST_ODONTOLOGO, enrollment)
	err := row.Scan(&enrollment)
	return err == nil
}

func (r *repository) HasTurno(id int) bool {
	rows, _ := r.db.Query(HAS_TURNO, id)

	if rows != nil {
		return true
	}
	return false
}