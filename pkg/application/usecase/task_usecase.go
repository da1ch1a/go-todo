package usecase

import (
	"da1ch1a/go-todo/pkg/domain/model"
	"da1ch1a/go-todo/pkg/domain/repository"
	"da1ch1a/go-todo/pkg/infra"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskUsecase struct {
	TaskRepository repository.TaskRepository
}

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (ts *TaskUsecase) List(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string][]Task{
		"tasks": {
			{
				ID: "1",
				Name: "test",
				Done: false,
				CreatedAt: "2024-01-01",
				UpdatedAt: "2024-01-01",
			},
			{
				ID: "2",
				Name: "test2",
				Done: false,
				CreatedAt: "2024-01-02",
				UpdatedAt: "2024-01-02",
			},
		},
	})
}

func (ts *TaskUsecase) Create(task model.Task) {
	ins, err := infra.Db.Prepare("INSERT INTO tasks (id, name, done, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
			log.Fatal(err)
	}
	res, err := ins.Exec(task.ID, task.Name, 0, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lastId)

	// _, err := infra.Db.Exec("INSERT INTO tasks id, name, done, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", task.ID, task.Name, 0, task.CreatedAt, task.UpdatedAt)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
