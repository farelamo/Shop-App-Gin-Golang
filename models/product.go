package models

import "time"

type Product struct {
	Id 				int 		`json:"id,omitempty"`
	Name			string		`json:"name,omitempty"`
	Stock 			int 		`json:"stock,omitempty"`
	Price 			int 		`json:"price,omitempty"`
	Description 	string 		`json:"description,omitempty"`
	CategoryId		int 		`json:"category_id,omitempty"`
	CreatedAt   	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   	time.Time 	`json:"updated_at,omitempty"`
}

type AddProduct struct {
	Name			string		`json:"name,omitempty"`
	Stock 			int 		`json:"stock,omitempty"`
	Price 			int 		`json:"price,omitempty"`
	Description 	string 		`json:"description,omitempty"`
	CategoryId		int 		`json:"category_id,omitempty"`
}