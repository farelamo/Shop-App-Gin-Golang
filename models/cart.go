package models

import "time"

type Cart struct {
	Id 				int 		`json:"id,omitempty"`
	UserId			int 		`json:"user_id,omitempty"`
	ProductId		int 		`json:"product_id,omitempty"`
	Paid 			bool 		`json:"paid,omitempty"`
	Checkout 		bool 		`json:"checkout,omitempty"`
	Total 			int 		`json:"total,omitempty"`
	CreatedAt   	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   	time.Time 	`json:"updated_at,omitempty"`
}

type AddCart struct {
	UserId			int 		`json:"user_id,omitempty"`
	ProductId		int 		`json:"product_id,omitempty"`
	Paid 			bool 		`json:"paid,omitempty"`
	Checkout 		bool 		`json:"checkout,omitempty"`
	Total 			int 		`json:"total,omitempty"`
}