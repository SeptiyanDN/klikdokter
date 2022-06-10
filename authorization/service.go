package authorization

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Services interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct {
}

var SECRET_KEY = []byte("KLIKDOKTER_DEVELOPMENT")

func NewServices() *jwtServices {
	return &jwtServices{}
}

func (s *jwtServices) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
