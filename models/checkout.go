package models

type CartCheckout struct {
	CartId 	 	[]int	`json:"cart_ids,omitempty"`
	BankName 	string  `json:"bank_name,omitempty"`
}

type Checkout struct {
	Id 		  		int   		`json:"id,omitempty"`
	UserId      	int			`json:"user_id,omitempty"`
	CartId 			int   		`json:"cart_id,omitempty"`
	BankName 		string   	`json:"bank_name,omitempty"`
	TotalCoupons 	float64   	`json:"total_coupons,omitempty"`
	TotalCheckout 	float64   	`json:"total_checkout,omitempty"`
}

type ResponseCheckout struct {
	Checkout 		[]Checkout	`json:"checkouts,omitempty"`
	TotalCoupons 	float64   	`json:"total_coupons,omitempty"`
	TotalCheckout 	float64   	`json:"total_checkout,omitempty"`
}