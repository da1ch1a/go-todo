package repositoryImpl

import (
	"da1ch1a/go-todo/pkg/domain/model"
	"database/sql"
	"log"
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
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Done, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) Create(task *model.Task) (error) {
	ins, err := r.db.Prepare("INSERT INTO tasks (id, name, done, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	res, err := ins.Exec(task.ID, task.Name, 0, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
