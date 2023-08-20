package ProductService

import "shop/models"

type ProductService interface {
	FindAll() (*[]models.Product, error)
	FindById(id int) (*models.Product, error)
	FindByCategory(id int) (*[]models.Product, error)
	Save(product *models.AddProduct) (*models.Product, error)
	Update(id int, product *models.Product) (int, error)
	Delete(id int) (int, error)
}