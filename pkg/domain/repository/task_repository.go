package repository

import (
	"da1ch1a/go-todo/pkg/domain/model"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *model.Task) error
	FindAll() ([]model.Task, error)
	Update(id uuid.UUID, done bool) error
}
