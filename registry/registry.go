package registry

import (
	"todo-app/api/handler"
	"todo-app/application/usecase"
	"todo-app/infrastructure/datasource"

	"go.uber.org/dig"
)

// アプリケーションで利用するDIコンテナを作成し、必要な依存関係を登録する
func BuildContainer() *dig.Container {
	container := dig.New()

	// datasource登録
	container.Provide(datasource.NewDatabase)
	container.Provide(datasource.NewUserRecordDatasource)

	// usecase登録
	container.Provide(usecase.NewUserRecordInteractor)

	//handler登録
	container.Provide(handler.NewUserHandler)
	return container
}
