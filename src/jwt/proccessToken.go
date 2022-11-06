package jwt

import (
	"errors"
	"strings"

	finder "modules/src/infrastructure/finder"
	models "modules/src/models"
	utils "modules/src/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserId string

func ProccessToken(token string) (*models.Claim, bool, string, error) {
	privateKey := []byte(utils.Config.Jwt.Secret)

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
		return nil, false, string(""), err
	}

	if !tkn.Valid {
		return nil, false, string(""), errors.New("invalod token")
	}

	_, exist, _ := finder.FindUserByEmail(claims.Email)

	if !exist {
		return nil, false, string(""), errors.New("invalid user")
	}

	Email = claims.Email
	UserId = claims.ID.Hex()

	return claims, exist, UserId, nil

}
