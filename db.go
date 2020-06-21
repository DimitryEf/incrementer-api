package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

type DbConnection struct {
	db *sql.DB
}

var (
	instance *DbConnection
	once     sync.Once
)

func NewDbConnection(config *Config) (*DbConnection, error) {

	var err error
	var db *sql.DB
	once.Do(func() {
		db, err = sql.Open("postgres", config.DBConnString)
		if err != nil {
			return
		}
		err = db.Ping()
		if err != nil {
			return
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS incrementer
			(
			num bigint,
			maximum_value bigint, 
			step_value bigint
			);`)
		if err != nil {
			return
		}

		row := db.QueryRow(`SELECT COUNT(*) FROM incrementer`)
		var count sql.NullInt64
		err = row.Scan(&count)

		if !count.Valid {
			_, err = db.Exec(`INSERT INTO incrementer (num, maximum_value, step_value)
			VALUES ($1, $2, $3)`, 0, MaximumInt, 1)
		}
		instance = &DbConnection{
			db: db,
		}
	})
	return instance, err
}
