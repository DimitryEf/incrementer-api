package container

import (
	"go.uber.org/dig"
	"os"
)

// Внедрение зависимостей с помощью dig
func BuildContainer() *dig.Container {
	// Создаем контейнер
	container := dig.New()

	// Устанавливаем зависимости
	exitIfErr(container.Provide(NewLogger))
	exitIfErr(container.Provide(NewConfig))
	exitIfErr(container.Provide(NewDbConnection))
	exitIfErr(container.Provide(NewPostgresRepo))
	exitIfErr(container.Provide(NewIncrementer))
	exitIfErr(container.Provide(NewGrpcApi))
	exitIfErr(container.Provide(NewServer))

	return container
}

// Функция завершения работы при ошибке
func exitIfErr(err error) {
	if err != nil {
		println(err)
		os.Exit(1)
	}
}

//Execute - проведение зависимостей и запуск сервера
func Execute() {
	container := BuildContainer()

	exitIfErr(container.Invoke(
		func(server *Server) {
			go server.Run()
			server.ReadyToStop() // Отслеживание сигналов OS для Graceful shutdown
		}))

	// DI вручную, без использования dig
	/*logger := NewLogger()
	config := NewConfig(logger)
	db, err := NewDbConnection(config)
	if err != nil {
		logger.Log.Fatal(err)
	}
	repo := NewPostgresRepo(db)
	inc := NewIncrementer(repo)
	api := NewGrpcApi(inc)
	server := NewServer(config, api)
	go server.Run()
	server.ReadyToStop()*/
}
