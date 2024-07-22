package usecase

import (
	"da1ch1a/go-todo/pkg/domain/model"
	"da1ch1a/go-todo/pkg/domain/repository"
	"da1ch1a/go-todo/pkg/infra"
	"fmt"
)

type TaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func (ts *TaskUsecase) List() []model.Task {
	rows, err := infra.Db.Query("SELECT id, name, done FROM tasks")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Done); err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, task)
	}

	return tasks
}
