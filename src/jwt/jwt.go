package jwt

import (
	"time"

	models "modules/src/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwt(user models.User) (string, error) {

	privateKey := []byte("apiresgolang")

	payload := jwt.MapClaims{
		"name":      user.Name,
		"lastName":  user.LastName,
		"dateBirth": user.DateBirth,
		"email":     user.Email,
		"avatar":    user.Avatar,
		"banner":    user.Banner,
		"Biografy":  user.Biografy,
		"Location":  user.Location,
		"WebSite":   user.WebSite,
		"id":        user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
