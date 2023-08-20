package CartService

import "shop/models"

type CartService interface {
	FindAll() (*[]models.Cart, error)
	FindById(id int) (*models.Cart, error)
	Save(cart *models.AddCart) (*models.Cart, error)
	Update(id int, cart *models.Cart) (int, error)
	Delete(id int) (int, error)
}