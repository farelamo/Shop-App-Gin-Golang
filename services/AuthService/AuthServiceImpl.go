package AuthService

import (
	"database/sql"
	"shop/models"
	"shop/utils/token"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	DB *sql.DB
}

func NewAuthService(DB *sql.DB) AuthService {
	return &AuthServiceImpl {
		DB: DB,
	}
}

func VerifyPassword(inputPassword string, passwordHashed *string) error {
	return bcrypt.CompareHashAndPassword([]byte(*passwordHashed), []byte(inputPassword))
}

func (a *AuthServiceImpl) LoginCheck(input *models.Login) (string, error){
	var user = models.User{}

	sql := `SELECT id,username,password FROM users WHERE username = ($1)`
	err := a.DB.QueryRow(sql, input.Username).Scan(&user.Id, &user.Username, &user.Pass)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(input.Pass, &user.Pass)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(&user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}