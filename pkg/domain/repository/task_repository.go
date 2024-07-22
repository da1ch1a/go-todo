package repository

import "da1ch1a/go-todo/pkg/domain/model"

type TaskRepository interface {
	FindAll() ([]model.Task, error)
}
