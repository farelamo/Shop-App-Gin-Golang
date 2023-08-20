package models

import "time"

type History struct {
	Id 				int 		`json:"id,omitempty"`
	UserId			int 		`json:"user_id,omitempty"`
	ProductId		int 		`json:"product_id,omitempty"`
	BankName 		string 		`json:"bank_name,omitempty"`
	CreatedAt   	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   	time.Time 	`json:"updated_at,omitempty"`
}

type AddHistory struct {
	UserId			int 		`json:"user_id,omitempty"`
	ProductId		int 		`json:"product_id,omitempty"`
	BankName 		string 		`json:"bank_name,omitempty"`
}