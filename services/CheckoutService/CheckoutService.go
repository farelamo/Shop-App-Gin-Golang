package CheckoutService

import "shop/models"

type CheckoutService interface {
	Save(userId int, cartCheckout *models.CartCheckout) (*models.ResponseCheckout, error)
}