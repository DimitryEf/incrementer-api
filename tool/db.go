package tool

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

// DbConnection - синглтон подключения к базе данных
type DbConnection struct {
	db *sql.DB
}

var (
	instance *DbConnection
	once     sync.Once // Гарантирует однократность вызова функции
)

func NewDbConnection(config *Config) (*DbConnection, error) {

	var err error
	var db *sql.DB
	once.Do(func() {
		// Настраиваем подключение к базе данных
		db, err = sql.Open("postgres", config.DBConnString)
		if err != nil {
			return
		}
		// Подключаемся к базе данных
		err = db.Ping()
		if err != nil {
			return
		}
		// Создаем таблицу incrementer, если она еще не существует
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS incrementer
			(
			num bigint,
			maximum_value bigint, 
			step_value bigint
			);`)
		if err != nil {
			return
		}

		// Проверяем наличие строки в таблице
		row := db.QueryRow(`SELECT COUNT(*) FROM incrementer`)
		var count int64
		err = row.Scan(&count)

		// Вставляем единственную строку с параметрами инкрементора по умолчанию
		if count == 0 {
			_, err = db.Exec(`INSERT INTO incrementer (num, maximum_value, step_value)
			VALUES ($1, $2, $3)`, 0, MaximumInt64, 1)
		}
		instance = &DbConnection{
			db: db,
		}
	})
	return instance, err
}
