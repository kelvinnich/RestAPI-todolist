package usecase

import (
	"authenctications/dto"
	"authenctications/model"
	"authenctications/repository"
	"log"

	"github.com/mashingan/smapping"
)


type TasksUseCase interface {
	AddTodoListUsecase(todoDTO dto.CreateTodoList) (*model.Tasks,error)
	UpdateTodoListUsecase(updateDTO dto.UpdateTodoList) (*model.Tasks, error)
	DeleteTodoListUsecase(id string)error
	GetAllTodoListUsecase() (*[]model.Tasks,error)
	GetTodoListByName(name string) (*model.Tasks, error)
}

type taskUseCase struct {
	taskRepo repository.TasksRepository
}

func(t *taskUseCase)AddTodoListUsecase(todoDTO dto.CreateTodoList) (*model.Tasks,error){
	var task model.Tasks
	err := smapping.FillStruct(&task, smapping.MapFields(&todoDTO))
	if err != nil {
		log.Printf("failed to map %v", err)
	}

	createTodo,err := t.taskRepo.AddTodoListRepository(task)
	if err != nil {
		log.Printf("failed to create todolist usecase %v", err)
	}

	return createTodo,nil
}

func (t *taskUseCase) UpdateTodoListUsecase(updateDTO dto.UpdateTodoList) (*model.Tasks, error) {
	task := model.Tasks{}
	if err := smapping.FillStruct(&task, smapping.MapFields(&updateDTO)); err != nil {
		log.Printf("failed to map %v", err)
		return nil, err
	}

	if updateTodo, err := t.taskRepo.UpdateTodoListRepository(updateDTO.Id,&task); err != nil {
		log.Printf("Failed to update todolist usecase %v", err)
		return nil, err
	} else {
		return updateTodo, nil
    } 
}


func(t *taskUseCase)DeleteTodoListUsecase(id string)error{
	err := t.taskRepo.DeleteTodoListRepository(id)
	if err != nil {
		log.Printf("failed to delete todolist usecase %v", err)
		return err
	}

	return nil
}

func(t *taskUseCase)GetAllTodoListUsecase() (*[]model.Tasks,error){
	todo,err := t.taskRepo.GetAllTodoListRepository()
	if err != nil {
		log.Printf("failed to get all data todolist usecase %v", err)
		return nil,err
	}

	return todo,nil
}

func(t *taskUseCase)GetTodoListByName(name string) (*model.Tasks,error){
	if todo,err := t.taskRepo.GetTodoListByName(name); err != nil{
		log.Printf("failed to get todolist by name %v", err)
		return nil,err
	}else {
		return todo,nil
	}
}

func NewTaskUsecase(taskRepository repository.TasksRepository) TasksUseCase{
	return &taskUseCase{taskRepo: taskRepository}
}