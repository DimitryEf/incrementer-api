package config

// start postgres server
// cd "D:\Program Files\PostgreSQL\11\bin"
// .\pg_ctl.exe -D "D:\Program Files\PostgreSQL\11\data" start

// Config - структура с конфигурацией сервиса
type Config struct {
	DBConnString    string
	Logger          *Logger
	Host            string
	Port            string
	HealthCheckPort string
}

func NewConfig(log *Logger) *Config {
	return &Config{
		//DBConnString:    "port=5432 host=localhost user=postgres password=soul dbname=postgres sslmode=disable",
		DBConnString:    "port=5432 host=db user=postgres password=mysecretpassword dbname=postgres sslmode=disable",
		Logger:          log,
		Host:            "",
		Port:            ":8080",
		HealthCheckPort: ":8081",
	}
}
