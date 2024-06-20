package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"clientes/models"
)

var secret = []byte("secret")
func VerifyToken(token string) (*models.UserClaims, bool) {
	payload, err  := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		return secret, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	if claims, ok := payload.Claims.(jwt.MapClaims); ok && payload.Valid {
		userClaims :=  models.UserClaims {
		
		}

		userClaims.Username = fmt.Sprint(claims["username"])
		tempRoleId := claims["roleid"].(float64)
		fmt.Println(tempRoleId)
		userClaims.RoleId = int(tempRoleId)

		return &userClaims, true
	}
	fmt.Println("Invalid Token")
	return nil, false
}