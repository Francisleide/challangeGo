package usecase

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/francisleide/ChallengeGo/gateways/db/repository"
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
	r repository.Repository
}

func NewAuthenticationUC(repo repository.Repository) AuthenticationUc {
	return AuthenticationUc{
		r: repo,
	}
}

func (a AuthenticationUc) Login(CPF, secret string) bool {
	account, ok := a.r.FindOne(CPF)
	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(secret))
	if err != nil {
		return false
	}
	return true

}

func (a AuthenticationUc) CreateToken(CPF string, secret string) (string, error) {
	b := a.Login(CPF, secret)
	if !b {
		log.Fatal("authentication error")
	}
	os.Setenv("ACCESS_SECRET", "asdhjkasjheee")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": CPF,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		log.Fatal("generate token error")
		return "", err
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
