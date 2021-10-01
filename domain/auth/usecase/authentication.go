package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/francisleide/ChallengeGo/domain/auth"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Claims struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}

type AuthenticationUc struct {
	r auth.Repository
}

func NewAuthenticationUC(repo auth.Repository) AuthenticationUc {
	return AuthenticationUc{
		r: repo,
	}
}

func (a AuthenticationUc) Login(CPF, secret string) error {
	account, err := a.r.FindOne(CPF)
	if err != nil {
		//TODO: add a sentinel
		return errors.New("invalid CPF")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(secret))
	if err != nil {
		//TODO: add a sentinel
		return errors.New("incorrect password")
	}
	return nil

}

func (a AuthenticationUc) CreateToken(CPF string, secret string) (string, error) {
	err := a.Login(CPF, secret)
	if err != nil {
		//TODO: add a senitinel
		return "", errors.New("invalid login")
	}
	os.Setenv("ACCESS_SECRET", "asdhjkasjheee")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": CPF,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		//TODO: add a sentinel
		return "", errors.New("it was not possible to generate the token")
	}

	return tokenString, nil
}

func Authentication(token *jwt.Token) interface{} {
	var accessUUID string
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		accessUUID, _ = claims["user"].(string)

	}
	return accessUUID

}
