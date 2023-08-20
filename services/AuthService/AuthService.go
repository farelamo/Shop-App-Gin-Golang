package AuthService

import "shop/models"

type AuthService interface {
	LoginCheck(input *models.Login) (string, error)
}