package CheckoutService

import (
	"database/sql"
	"errors"
	"math"
	"shop/models"
	cart "shop/services/CartService"

	"github.com/lib/pq"
	"fmt"
)

type CheckoutServiceImpl struct {
	DB *sql.DB
}

func NewCheckoutService(DB *sql.DB) CheckoutService {
	return &CheckoutServiceImpl{
		DB: DB,
	}
}

func (c *CheckoutServiceImpl) Save(userId int, cartCheckout *models.CartCheckout) (*models.ResponseCheckout, error) {
	var Results models.ResponseCheckout
	var CartCheckouts []models.Checkout

	cartService := cart.NewCartService(c.DB)

	sql 	  := `SELECT * FROM carts WHERE id = ANY($1) AND checkout = false`
	rows, err := c.DB.Query(sql, pq.Array(cartCheckout.CartId))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	coupon := 0
	var totalCheckout float64
	for rows.Next() {
		var Cart models.Cart
		Checkout := models.Checkout{}

		err := rows.Scan(&Cart.Id, &Cart.UserId, &Cart.ProductId, &Cart.Checkout, &Cart.Amount, &Cart.Total, &Cart.CreatedAt, &Cart.UpdatedAt)
		if err != nil {
			return nil, err
		}

		Checkout.CartId   	   = Cart.Id
		Checkout.UserId   	   = userId
		Checkout.BankName 	   = cartCheckout.BankName
		Checkout.TotalCheckout = float64(Cart.Total)
		Checkout.TotalCoupons  = 0

		if Cart.Total > 50000 {
			coupon += 1
			totalCheckout += float64(Cart.Total)
			Checkout.TotalCoupons += 1
		}

		Cart.Checkout = true
		_, err = cartService.Update(Cart.Id, &Cart)
		if err != nil {
			return nil, err
		}

		sql := `INSERT INTO checkouts (cart_id, user_id, bank_name, total_coupons, total_checkout) VALUES ($1, $2, $3, $4, $5) Returning *`
		err = c.DB.QueryRow(sql, Checkout.CartId, Checkout.UserId, Checkout.BankName, Checkout.TotalCoupons, Checkout.TotalCheckout,).Scan()

		CartCheckouts = append(CartCheckouts, Checkout)
	}

	fmt.Println(len(CartCheckouts))

	if len(CartCheckouts) <= 0 {
		return nil, errors.New("Data has already checkout")
	}

	calculate := math.Floor(totalCheckout / 100000) 
	if calculate > 0 {
		coupon += int(calculate)
	}

	Results.Checkout 	  = CartCheckouts
	Results.TotalCheckout = totalCheckout
	Results.TotalCoupons  = float64(coupon)

	return &Results, nil
}