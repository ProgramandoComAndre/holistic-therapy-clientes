package middleware

import (
	"clientes/models"
	"clientes/utils"
	"fmt"
	"net/http"
	"strings"
)


func VerifyAuth(r *http.Request) (*models.UserClaims, bool) {
	tokenString := r.Header.Get("Authorization")
	
	if tokenString == "" {
		return nil, false
	}

	parts := strings.Split(tokenString, " ")
	if(len(parts) != 2) {
		return nil, false
	}
	fmt.Println(parts[1])
	userClaims, success := utils.VerifyToken(parts[1])
	return userClaims, success
}
