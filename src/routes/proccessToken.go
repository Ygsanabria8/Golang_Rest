package routes

import (
	"errors"
	"strings"

	dataBase "modules/src/dataBase"
	models "modules/src/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserId string

func ProccessToken(token string) (*models.Claim, bool, string, error) {
	privateKey := []byte("apiresgolang")

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalod token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return claims, false, string(""), err
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalod token")
	}

	_, exist, _ := dataBase.FindUserByEmail(claims.Email)
	if !exist {
		return claims, false, string(""), errors.New("invalid user")
	}

	Email = claims.Email
	UserId = claims.ID.Hex()

	return claims, exist, UserId, nil

}
