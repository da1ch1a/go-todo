package repository

import "da1ch1a/go-todo/pkg/domain/model"

type TaskRepository interface {
	Create(task *model.Task) error
	FindAll() ([]model.Task, error)
}
