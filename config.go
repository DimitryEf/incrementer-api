package main

import "time"

// start postgres server
// cd "D:\Program Files\PostgreSQL\11\bin"
// .\pg_ctl.exe -D "D:\Program Files\PostgreSQL\11\data" start

type Config struct {
	DBConnString    string
	Log             *Logger
	Host            string
	Port            string
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxHeaderBytes  int
}

func NewConfig(log *Logger) *Config {
	return &Config{
		DBConnString:    "port=5432 host=localhost user=postgres password=soul dbname=postgres sslmode=disable",
		Log:             log,
		Host:            "",
		Port:            ":8080",
		ShutdownTimeout: 10 * time.Second,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		MaxHeaderBytes:  1 << 20,
	}
}
