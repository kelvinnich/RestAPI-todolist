package repository

import (
	"authenctications/model"
	"database/sql"
	"errors"
	"log"
)


type TasksRepository interface {
	AddTodoListRepository(t model.Tasks) (*model.Tasks, error)
	UpdateTodoListRepository(id string, t model.Tasks) (*model.Tasks, error)
	DeleteTodoListRepository(id string) error
	GetAllTodoListRepository() (*[]model.Tasks, error)
}

type taskConnections struct {
	db *sql.DB
}

func(db *taskConnections)AddTodoListRepository(t model.Tasks) (*model.Tasks, error){
	query := `INSERT INTO tasks(name, description, deadline, status, users_id) VALUES($1, $2, $3, $4, $5)`

	_,err := db.db.Exec(query, t.Name, t.Description,t.Deadline, t.Status, t.Users_id)
	if err != nil {
		log.Printf("failed to create todolist repository %v", err)
		return nil,err
	}

	return &t, nil
}

func(db *taskConnections) UpdateTodoListRepository(id string, t model.Tasks) (*model.Tasks, error){
	query := `UPDATE tasks SET name=$1, description=$2, deadline=$3, status=$4, users_id=$5 WHERE id=$6`

	_,err := db.db.Exec(query, t.Name, t.Description, t.Deadline, t.Status, t.Users_id, id)
	if err != nil {
		log.Printf("failed to update todolist repository %v", err)
		return nil,err
	}

	return &t, nil
}

func(db *taskConnections) DeleteTodoListRepository(id string) error{
	query := `DELETE FROM users WHERE id = $1`

	_,err := db.db.Exec(query, id)
	if err != nil {
		log.Printf("failed to delete todolist repository %v", err)
		return err
	}

	return nil
}

func(db *taskConnections)GetAllTodoListRepository() (*[]model.Tasks, error){
	rows,err := db.db.Query("SELECT id,name,description,deadline,status,users_id FROM tasks")
	if err != nil {
		log.Printf("failed to get all data todolist repository %v", err)
	}
	
	defer rows.Close()

	tasks := make([]model.Tasks, 0)

	for rows.Next() {
		var task model.Tasks
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Deadline, &task.Status, &task.Users_id)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil,err
	}

	if len(tasks) == 0 {
		return nil, errors.New("task not found")
	}

	return &tasks, nil
}

func NewTaskRepository(db *sql.DB) TasksRepository{
	return &taskConnections{db: db}
}