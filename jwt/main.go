package main

import (
	"crypto/hmac"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"

	"github.com/golang-jwt/jwt"
)

func main() {
	slog.Info("JWT Tester")
	hmacSampleSecret := []byte("my_secret_key"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return hmacSampleSecret, nil
}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
if err != nil {
	slog.Error(err.Error())
}

if claims, ok := token.Claims.(jwt.MapClaims); ok {
	fmt.Println(claims["foo"], claims["nbf"])
} else {
	fmt.Println(err)
}

}
