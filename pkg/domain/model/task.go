package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type Task struct {
	ID        string
	Name      string
	Done      bool
	CreatedAt string
	UpdatedAt string
}

func NewTaskFromJson(jsonBody map[string]interface{}) (*Task, error) {
	name, ok := jsonBody["name"].(string)
	if !ok {
		return nil, xerrors.New("name is required")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Task{
		ID: uuid.String(),
		Name: name,
		Done: false,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
