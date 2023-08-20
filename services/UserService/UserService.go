package UserService

import "shop/models"

type UserService interface {
	FindAll() (*[]models.User, error)
	FindById(id int) (*models.User, error)
	Save(user *models.AddUser) (*models.User, error)
	Update(id int, user *models.User) (int, error)
	Delete(id int) (int, error)
}