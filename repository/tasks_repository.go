package repository

import (
	"authenctications/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)
type TasksRepository interface {
	AddTodoListRepository(t model.Tasks) (*model.Tasks, error)
	UpdateTodoListRepository(id string, t *model.Tasks) (*model.Tasks, error)
	DeleteTodoListRepository(id string) error
	GetAllTodoListRepository() (*[]model.Tasks, error)
	GetTodoListByName(name string) (*model.Tasks, error)
}

type taskConnections struct {
	db *sql.DB
}

func(db *taskConnections)AddTodoListRepository(t model.Tasks) (*model.Tasks, error){
	query := `INSERT INTO tasks(id, name, description, deadline, status, users_id) VALUES($1, $2, $3, $4, $5, $6)`

	_,err := db.db.Exec(query ,t.Id,t.Name, t.Description,t.Deadline, t.Status, t.Users_id)
	if err != nil {
		log.Printf("failed to create todolist repository %v", err)
		return nil,err
	}
	return &t, nil
}

func (db *taskConnections) UpdateTodoListRepository(id string, t *model.Tasks) (*model.Tasks, error) {
	var idExists bool
	err := db.db.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE id=$1)", id).Scan(&idExists)
	if err != nil {
			log.Printf("failed to check if ID exists in database: %v", err)
			return nil, err
	}
	if !idExists {
			return nil, fmt.Errorf("ID %s not found in database",id)
	}

	query := `UPDATE tasks SET name=$1, description=$2, deadline=$3, status=$4, users_id=$5 WHERE id=$6`
	result, err := db.db.Exec(query, t.Name, t.Description, t.Deadline, t.Status, t.Users_id, id)
	if err != nil {
			log.Printf("failed to update todolist repository: %v", err)
			return nil, err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
			log.Printf("failed to retrieve number of affected rows: %v", err)
			return nil, err
	} else if rowsAffected == 0 {
			return nil, errors.New("no rows affected")
	}
	return t, nil
}

func(db *taskConnections) DeleteTodoListRepository(id string) error{
	query := `DELETE FROM tasks WHERE id = $1`

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
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Deadline, &task.Status, &task.Users_id)
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

func (db *taskConnections) GetTodoListByName(name string) (*model.Tasks, error) {
	var task model.Tasks
	query := `SELECT id,name,description,deadline,status,users_id FROM tasks WHERE name = $1`

	err := db.db.QueryRow(query, name).Scan(&task.Id, &task.Name, &task.Description, &task.Deadline, &task.Status, &task.Users_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no todo list found with the given name")
		} else {
			log.Printf("failed get todolist by name repository %v", err)
			return nil, err
		}
	}

	return &task, nil
}


func NewTaskRepository(db *sql.DB) TasksRepository{
	return &taskConnections{db: db}
}