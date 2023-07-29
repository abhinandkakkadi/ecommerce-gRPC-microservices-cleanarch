package helper

import (
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



