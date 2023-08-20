package models

import "time"

type Category struct {
	Id 				int 		`json:"id,omitempty"`
	Name			string		`json:"name,omitempty"`
	CreatedAt   	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   	time.Time 	`json:"updated_at,omitempty"`
}

type AddCategory struct {
	Name   string	`json:"name,omitempty"`
}