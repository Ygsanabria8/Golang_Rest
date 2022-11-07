package jwt

import (
	"time"

	models "modules/src/models"

	utils "modules/src/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwt(user *models.User) (string, error) {

	privateKey := []byte(utils.Config.Jwt.Secret)

	payload := jwt.MapClaims{
		"name":     user.Name,
		"lastName": user.LastName,
		"email":    user.Email,
		"id":       user.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
