package registry

import (
	"da1ch1a/go-todo/pkg/application/usecase"
	"da1ch1a/go-todo/pkg/domain/repository"
	"da1ch1a/go-todo/pkg/infra/repositoryImpl"
	"database/sql"
)

type Registry struct {
	Db *sql.DB
	taskRepository repository.TaskRepository

	TaskUsecase usecase.TaskUsecase
}

func BuildRegistry(db *sql.DB) *Registry {
	// リポジトリの初期化
	taskRepository := repositoryImpl.NewTaskRepositoryImpl(db)

	// ユースケースの初期化
	taskUsecase := usecase.TaskUsecase{
		TaskRepository: taskRepository,
	}

	registry := &Registry{
		Db: db,
		taskRepository: taskRepository,

		TaskUsecase: taskUsecase,
	}

	return registry
}
