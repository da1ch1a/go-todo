package registry

import (
	"da1ch1a/go-todo/pkg/domain/repository"
	"da1ch1a/go-todo/pkg/infra/repositoryImpl"
	"database/sql"
)

type Registry struct {
	Db *sql.DB
	taskRepository repository.TaskRepository
}

func BuildRegistry(db *sql.DB) *Registry {
	// リポジトリの初期化
	taskRepository := repositoryImpl.NewTaskRepositoryImpl(db)

	registry := &Registry{
		// 今は不要だが、他のリポジトリが増えたときのためにdbをフィールドに持つ
		Db: db,
		taskRepository: taskRepository,
	}

	return registry
}
