package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type Claim struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func VerifyToken(r *http.Request, lenguaje string, key string) (string, int) {
	var description string
	var code int
	var publicRsa = os.Getenv("PATH_JWT") + key
	tokenString := ExtractToken(r)
	publicBytes, err := ioutil.ReadFile(publicRsa)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			description, code = ReturnTokenError(r, vErr, lenguaje, true)
			return description, code
		default:
			description = Language(lenguaje, "INVALID_TOKEN")
			code = http.StatusUnauthorized
			return description, code
		}
	}
	if token != nil {
		code = http.StatusOK
	}

	return description, code
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractIdUser(tokenStr string) string {
	path := os.Getenv("PATH_JWT")
	var publicRsa = path + os.Getenv("PUBLIC_RSA")
	publicBytes, err := ioutil.ReadFile(publicRsa)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(err)
	return claims["user"].(string)

}

func ReturnTokenError(r *http.Request, vErr *jwt.ValidationError, lenguaje string, verifyAgain bool) (string, int) {
	var description string
	var code int
	switch vErr.Errors {
	case jwt.ValidationErrorExpired:
		description = Language(lenguaje, "TOKEN_EXPIRED")
		code = http.StatusNonAuthoritativeInfo
		return description, code
	case jwt.ValidationErrorSignatureInvalid:
		//Si es invalida, se realiza verificación con la firma de invitado
		return VerifyAgain(r, lenguaje, verifyAgain)
	default:
		description = Language(lenguaje, "INVALID_TOKEN")
		code = http.StatusUnauthorized
		return description, code
	}
}

func VerifyAgain(r *http.Request, lenguaje string, verifyAgain bool) (string, int) {
	var description string
	var code int
	if verifyAgain {
		description, code = VerifyToken(r, lenguaje, os.Getenv("PUBLIC_RSA_GUEST"))
		return description, code
	}
	description = Language(lenguaje, "SIGNATURE_MATCH")
	code = http.StatusUnauthorized
	return description, code
}
