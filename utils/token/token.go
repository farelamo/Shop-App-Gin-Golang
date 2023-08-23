package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId *int) (string, error) {

	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	expired_time, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRED"))
	if err != nil {
		return "", err
	}


	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] 	 = *userId
	claims["exp"]		 = time.Now().Add(time.Hour * time.Duration(expired_time)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func TokenValidation(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err 		:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	splitToken  := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}

func GetUserId(c *gin.Context) (int, error) {
	tokenString := ExtractToken(c)
	token, err  := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	data   := claims["user_id"]
	
	return int(data.(float64)), nil
}