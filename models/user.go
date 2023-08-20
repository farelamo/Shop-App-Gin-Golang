package models

import "time"

type User struct {
	Id 				int 		`json:"id,omitempty"`
	Name 			string 		`json:"name,omitempty"`
	Username 		string 		`json:"username,omitempty"`
	Pass 			string 		`json:"password,omitempty"`
	Age 			int 		`json:"age,omitempty"`
	CreatedAt   	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   	time.Time 	`json:"updated_at,omitempty"`
}

type AddUser struct {
	Name 			string 		`json:"name,omitempty"`
	Username 		string 		`json:"username,omitempty"`
	Pass 			string 		`json:"password,omitempty"`
	Age 			int 		`json:"age,omitempty"`
}

type Login struct {
	Username 		string 		`json:"username,omitempty"`
	Pass 			string 		`json:"password,omitempty"`
}