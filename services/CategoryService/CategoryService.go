package CategoryService

import "shop/models"

type CategoryService interface {
	FindAll() (*[]models.Category, error)
	FindById(id int) (*models.Category, error)
	Save(category *models.AddCategory) (*models.Category, error)
	Update(id int, cart *models.Category) (int, error)
	Delete(id int) (int, error)
}