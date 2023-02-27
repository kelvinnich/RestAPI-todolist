package controller

import (
	"authenctications/dto"
	"authenctications/middleware"
	"authenctications/usecase"
	"authenctications/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	CreateTodoList(c *gin.Context)
	UpdateTodoList(c *gin.Context)
	DeleteTodoList(c *gin.Context)
	GetAllTodoList(c *gin.Context)
}

type taskController struct {
	taskUsecase usecase.TasksUseCase
	jwt usecase.JwtUseCase
	r *gin.Engine
}

func(t *taskController)CreateTodoList(c *gin.Context){
	var task dto.CreateTodoList
	task.Id = util.NewUUID()
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	todo,err := t.taskUsecase.AddTodoListUsecase(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "failed to create todolist" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func(t *taskController)UpdateTodoList(c *gin.Context){
	var task dto.UpdateTodoList
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	todo,err := t.taskUsecase.UpdateTodoListUsecase(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "failed to update todolist: " + err.Error()}) 
		return
	}
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo}) 
}



func(t *taskController)DeleteTodoList(c *gin.Context){
	id := c.Param("id")


	err := t.taskUsecase.DeleteTodoListUsecase(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "failed to delete todolist" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "success delete data"})
}

func(t *taskController)GetAllTodoList(c *gin.Context){
	todolist,err := t.taskUsecase.GetAllTodoListUsecase()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "failed to get all data todolist"+err.Error()})
		return
	}

	c.JSON(http.StatusOK, todolist)
}

func NewTodoListController(taskUsecase usecase.TasksUseCase, jwt usecase.JwtUseCase, r *gin.Engine) TaskController{
	todolist := taskController{
		taskUsecase: taskUsecase,
		jwt: jwt,
		r: r,
	}

	todo := r.Group("/todolist", middleware.Authorize(jwt))
	{
		todo.POST("/addTodolist", todolist.CreateTodoList)
		todo.PUT("/updateTodo/:id", todolist.UpdateTodoList)
		todo.DELETE("/:id", todolist.DeleteTodoList)
		todo.GET("/", todolist.GetAllTodoList)
	}

	return &todolist
}