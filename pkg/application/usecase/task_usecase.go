package usecase

import (
	"da1ch1a/go-todo/pkg/domain/model"
	"da1ch1a/go-todo/pkg/domain/repository"
	"encoding/json"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

type TaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func (u *TaskUsecase) List() []model.Task {
	tasks, err := u.TaskRepository.FindAll()
	if err != nil {
		log.Printf("failed to find all tasks: %#v", err)
		return nil
	}

	return tasks
}

func (u *TaskUsecase) Create(c echo.Context) error {
	jsonBody, err := u.getBody(c)
	if err != nil {
		return err
	}

	task, err := model.NewTaskFromJson(jsonBody)
	if err != nil {
		return err
	}

	err = u.TaskRepository.Create(task)
	if err != nil {
		return err
	}

	return nil
}

func (u *TaskUsecase) getBody(c echo.Context) (map[string]interface{}, error) {
	body := c.Request().Body
	defer func() {
		_ = body.Close()
	}()
	data, err := io.ReadAll(body)
	if err != nil {
		log.Printf("failed to read body: %#v", err)
		return nil, err
	}

	// parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(data, &jsonBody)
	if err != nil {
		log.Printf("failed to unmarshal json: %#v", err)
		return nil, err
	}

	return jsonBody, nil
}
