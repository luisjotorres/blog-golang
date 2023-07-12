package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// ToDO obtener esto de un env file
var jwtKey = []byte("my_secret_key")

type Claims struct {
	UserId           uint   `json:"author_id"`
	UserType         string `json:"type"`
	IpAddress        string `json:"ip_address"`
	TimeToStartVisit time.Time
	jwt.RegisteredClaims
}

/*
Construye un token utilizando JWT
Este token sera utilizado para todas las transacciones de la app

Si es de tipo author, utilizaremos para validar sus permisos de escritura
Si es de tipo visit, lo utilizaremos para seguir las acciones del usuario
*/
func JwtBuilder(userType, ipAddress string, userId uint) string {
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := Claims{
		UserId:    userId,
		UserType:  userType,
		IpAddress: ipAddress,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}
