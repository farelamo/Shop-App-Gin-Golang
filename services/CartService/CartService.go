package CartService

import "shop/models"

type CartService interface {
	FindAll(userId int) (*[]models.Cart, error)
	FindById(id int) (*models.Cart, error)
	Save(userId int, cart *models.AddCart) (*models.Cart, error)
	Update(id int, cart *models.Cart) (int, error)
	Delete(userId int, id int) (int, error)
}