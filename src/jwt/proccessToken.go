package jwt

import (
	"errors"
	"strings"

	models "modules/src/models"
	utils "modules/src/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

func ProccessToken(token string) (bool, error) {
	privateKey := []byte(utils.Config.Jwt.Secret)

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return false, errors.New("invalod token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return false, err
	}

	if !tkn.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil

}
