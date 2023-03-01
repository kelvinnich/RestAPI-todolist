package model

import "time"

type Tasks struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Deadline time.Time `json:"deadline"`
	Status bool `json:"status"`
	Users_id string `json:"users_id"`
}