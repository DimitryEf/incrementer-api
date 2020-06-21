package repo

import (
	"database/sql"
	"errors"
	"github.com/DimitryEf/incrementer-api/model"
	"github.com/DimitryEf/incrementer-api/tool"
)

// PostgresRepo - репозиторий базы данных PostgresQL
type PostgresRepo struct {
	DB *sql.DB
}

// Ошибки при получении пустых значений
var (
	ErrNumberIsNull       = errors.New("increment value is null in storage")
	ErrMaximumValueIsNull = errors.New("maximum value is null in storage")
	ErrStepValueIsNull    = errors.New("step value is null in storage")
)

// Конструктор репозитория
func NewPostgresRepo(dbConn *tool.DbConnection) *PostgresRepo {
	return &PostgresRepo{
		DB: dbConn.DB,
	}
}

// Методы, имплементирующие интерфейс Repo
func (repo *PostgresRepo) GetNumber() (int64, error) {
	query := "SELECT num FROM incrementer"
	row := repo.DB.QueryRow(query)
	var num sql.NullInt64
	if err := row.Scan(&num); err != nil {
		return 0, err
	}
	if !num.Valid {
		return 0, ErrNumberIsNull
	}
	return num.Int64, nil
}

func (repo *PostgresRepo) SetNumber(num int64) error {
	query := "UPDATE incrementer SET num = $1"
	_, err := repo.DB.Exec(query, num)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) GetParams() (model.Params, error) {
	query := "SELECT num, maximum_value, step_value FROM incrementer"
	row := repo.DB.QueryRow(query)
	var num, max, step sql.NullInt64
	if err := row.Scan(&num, &max, &step); err != nil {
		return model.Params{}, err
	}
	if !num.Valid {
		return model.Params{}, ErrNumberIsNull
	}
	if !max.Valid {
		return model.Params{}, ErrMaximumValueIsNull
	}
	if !step.Valid {
		return model.Params{}, ErrStepValueIsNull
	}

	return model.Params{
		Number:       num.Int64,
		MaximumValue: max.Int64,
		StepValue:    step.Int64,
	}, nil
}

func (repo *PostgresRepo) SetMaximumValue(maximumValue int64) error {
	query := "UPDATE incrementer SET maximum_value = $1"
	_, err := repo.DB.Exec(query, maximumValue)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) SetStepValue(stepValue int64) error {
	query := "UPDATE incrementer SET step_value = $1"
	_, err := repo.DB.Exec(query, stepValue)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) SetParams(maximumValue, stepValue int64) error {
	query := "UPDATE incrementer SET maximum_value = $1, step_value = $2"
	_, err := repo.DB.Exec(query, maximumValue, stepValue)
	if err != nil {
		return err
	}
	return nil

}

func (repo *PostgresRepo) SetMaximumValueAndZeroNumber(maximumValue int64) error {
	query := "UPDATE incrementer SET maximum_value = $1, num = $2"
	_, err := repo.DB.Exec(query, maximumValue, 0)
	if err != nil {
		return err
	}
	return nil
}
