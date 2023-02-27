package dto

import (
	"time"
)

type CreateTodoList struct {
	Id string `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Deadline time.Time `json:"deadline" form:"deadline" binding:"required"`
	Status bool `json:"status" form:"status"`
	Users_id string `json:"users_id" form:"users_id" binding:"required"`
}

type UpdateTodoList struct {
	Id string `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Deadline time.Time `json:"deadline" form:"deadline" binding:"required"`
	Status bool `json:"status" form:"status"`
	Users_id string `json:"users_id" form:"users_id" binding:"required"`
}