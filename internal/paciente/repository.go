package Paciente

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DaianaServin/ProyectoApi/internal/domain"

)

type Repository interface {
	Save(ctx context.Context, p domain.Paciente) (int, error)
	GetByID(id int) (domain.Paciente, error)
	Update(p domain.Paciente) (domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Exists(DNI string) bool
	HasShifts(id int) bool
}

const (
	SAVE_PACIENTE         = "INSERT INTO pacientes(name, surname, address,DNI,fechaTurno) VALUES (?,?,?,?,?);"
	GET_PACIENTE_BY_ID    = "SELECT * FROM pacientes WHERE id = ?;"
	UPDATE_PACIENTE       = "UPDATE pacientes SET name = ?, surname = ?, address = ?, DNI = ?, fechaTurno = ? WHERE id = ?;"
	DELETE_PACIENTE_BY_ID = "DELETE FROM pacientes WHERE id = ?;"
	EXIST_PACIENTE        = "SELECT DNI FROM pacientes WHERE DNI = ?"
	HAS_TURNO            = "SELECT p.* FROM pacientes p INNER JOIN turno t ON p.id = t.paciente_id WHERE paciente_id = ?;"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, p domain.Paciente) (int, error) {
	stmt, err := r.db.Prepare(SAVE_PACIENTE)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.Name, &p.Surname, &p.Address, &p.DNI, &p.TurnoFecha)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	row := r.db.QueryRow(GET_PACIENTE_BY_ID, id)
	p := domain.Paciente{}
	err := row.Scan(&p.ID, &p.Name, &p.Surname, &p.Address, &p.DNI, &p.TurnoFecha)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (r *repository) Update(p domain.Paciente) (domain.Paciente, error) {
	stmt, err := r.db.Prepare(UPDATE_PACIENTE)
	if err != nil {
		return domain.Paciente{}, err
	}

	res, err := stmt.Exec(&p.Name, &p.Surname, &p.Address, &p.DNI, &p.DischargeDate, &p.ID)
	if err != nil {
		return domain.Paciente{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Paciente{}, err
	}

	return p, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DELETE_PACIENTE_BY_ID)
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

func (r *repository) Exists(DNI string) bool {
	row := r.db.QueryRow(EXIST_PACIENTE, DNI)
	err := row.Scan(&DNI)
	return err == nil
}

func (r *repository) HasTurnos(id int) bool {
	rows, _ := r.db.Query(HAS_TURNO, id)

	if rows != nil {
		return true
	}
	return false
}