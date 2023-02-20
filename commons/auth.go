package commons

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetAuthToken(username, password, foundPassword string) (string, *RequestError) {

	expirationTime := time.Now().Add(time.Minute * 15)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		ErrorLogger.Println(err.Error())
		return "", &RequestError{StatusCode: http.StatusUnauthorized, ErrorOccurredIn: "auth GetAuthToken", Err: err.Error()}
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(password))
	if err != nil {
		WarningLogger.Println(err.Error())
		return "", &RequestError{StatusCode: http.StatusUnauthorized, ErrorOccurredIn: "auth GetAuthToken", Err: err.Error()}
	} else {
		return tokenString, nil
	}
}

func IsAuthorized(tokenString string) (bool, Claims, *RequestError) {

	splitToken := strings.Split(tokenString, "Bearer ")
	reqToken := splitToken[1]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(reqToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ErrorLogger.Println(err.Error())
			return false, *claims, &RequestError{StatusCode: http.StatusUnauthorized, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
		}
		return false, *claims, &RequestError{StatusCode: http.StatusUnauthorized, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
	}

	if !tkn.Valid {
		return false, *claims, &RequestError{StatusCode: http.StatusUnauthorized, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
	}

	return true, *claims, nil

}

func RefreshToken(tokenString string) (string, *RequestError) {

	splitToken := strings.Split(tokenString, "Bearer ")
	reqToken := splitToken[1]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(reqToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ErrorLogger.Println(err.Error())
			return "", &RequestError{StatusCode: http.StatusUnauthorized, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
		}
		return "", &RequestError{StatusCode: http.StatusUnauthorized, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
	}

	if !tkn.Valid {
		return "", &RequestError{StatusCode: http.StatusUnauthorized, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newReqToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", &RequestError{StatusCode: http.StatusInternalServerError, Err: err.Error(), ErrorOccurredIn: "auth IsAuthorized"}
	}

	return newReqToken, nil

}
