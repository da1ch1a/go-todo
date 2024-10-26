package repositoryImpl

import (
	"context"
	"da1ch1a/go-todo/pkg/domain/model"
	"da1ch1a/go-todo/schemas"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (r *TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	schemasTasks, err := schemas.Tasks().All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}


	tasks := make([]model.Task, len(schemasTasks))
	for i, task := range schemasTasks {
		tasks[i] = model.Task{
			ID:   task.ID,
			Name: task.Name,
			Done:      task.Done == 1,
			CreatedAt: task.CreatedAt.String(),
			UpdatedAt: task.UpdatedAt.String(),
		}
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) Create(task *model.Task) error {
	done := uint8(0)
	if task.Done {
		done = 1
	}
	t := schemas.Task{
		ID:   task.ID,
		Name: task.Name,
		Done: done,
	}

	err := t.Insert(context.Background(), r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
