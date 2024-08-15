package main

import (
	"todo-app/api"
	"todo-app/core/logger"
	"todo-app/infrastructure/datasource"
)

func main() {
	logger.Initialize()
	datasource.Migrate("file://./_migrations")
	api.Run()
}
