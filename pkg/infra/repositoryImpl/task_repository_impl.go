package repositoryImpl

import (
	"da1ch1a/go-todo/pkg/domain/model"
	"database/sql"
)

// 名前空間
type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepositoryImpl(db *sql.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{
		db: db,
	}
}

func (tr *TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	return nil, nil
}
