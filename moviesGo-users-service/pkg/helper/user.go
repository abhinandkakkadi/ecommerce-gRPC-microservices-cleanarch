package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/utils/models"
	"github.com/golang-jwt/jwt"
)

type authCustomClaimsUsers struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func init() {

}

func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {

	claims := &authCustomClaimsUsers{
		Id:    userID,
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("132457689"))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GenerateAccessToken(user models.UserDetailsResponse) (string, error) {

	expirationTime := time.Now().Add(30 * time.Minute)
	tokenString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}


func ValidateUser(tokenString string) (int,error) {

	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsUsers{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte("132457689"), nil // Replace with your actual secret key
	})

	if err != nil {
		return 0,errors.New("invalid signing method")
	}

	claims, ok := token.Claims.(*authCustomClaimsUsers)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.Id,nil
	
}

