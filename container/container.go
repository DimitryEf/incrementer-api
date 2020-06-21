package container

import (
	"github.com/DimitryEf/incrementer-api/api"
	"github.com/DimitryEf/incrementer-api/config"
	"github.com/DimitryEf/incrementer-api/repo"
	"github.com/DimitryEf/incrementer-api/server"
	"github.com/DimitryEf/incrementer-api/tool"
	"github.com/DimitryEf/incrementer-api/usecase"
	"go.uber.org/dig"
	"os"
)

// Внедрение зависимостей с помощью dig
func BuildContainer() *dig.Container {
	// Создаем контейнер
	container := dig.New()

	// Устанавливаем зависимости
	exitIfErr(container.Provide(config.NewLogger))
	exitIfErr(container.Provide(config.NewConfig))
	exitIfErr(container.Provide(tool.NewDbConnection))
	exitIfErr(container.Provide(repo.NewPostgresRepo))
	exitIfErr(container.Provide(usecase.NewIncrementer))
	exitIfErr(container.Provide(api.NewGrpcApi))
	exitIfErr(container.Provide(server.NewServer))

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
	/*container := BuildContainer()

	exitIfErr(container.Invoke(
		func(server *server.Server) {
			go server.Run()
			server.ReadyToStop() // Отслеживание сигналов OS для Graceful shutdown
		}))*/

	// DI вручную, без использования dig
	logger := config.NewLogger()
	config := config.NewConfig(logger)
	db, err := tool.NewDbConnection(config)
	if err != nil {
		logger.Log.Fatal(err)
	}
	repo := repo.NewPostgresRepo(db)
	inc := usecase.NewIncrementer(repo)
	api := api.NewGrpcApi(inc)
	server := server.NewServer(config, api)
	go server.Run()
	server.ReadyToStop()
}
