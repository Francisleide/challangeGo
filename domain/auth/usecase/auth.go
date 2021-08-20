package usecase

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/gateways/db/repository"
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

type AuthUc struct {
	r repository.Repository
}

func NewAuth(repo repository.Repository) AuthUc {
	return AuthUc{
		r: repo,
	}
}

func (a AuthUc) Login(CPF, secret string) bool {
	var account entities.Account
	account.CPF = CPF
	account.Secret = secret
	acc := a.r.FindOne(account.CPF)
	if reflect.DeepEqual(acc, entities.Account{}) {
		fmt.Printf("Usuário/senha incorreto")
		return false
	}
	fmt.Println("Senha cript: ", acc.Secret)
	fmt.Println("Senha não cript: ", secret)
	err := bcrypt.CompareHashAndPassword([]byte(acc.Secret), []byte(secret))
	if err != nil {
		fmt.Println("Usuário/senha incorreto")
		return false
	}
	return true

}

func (a AuthUc) CreateToken(CPF string, secret string) (string, error) {
	b := a.Login(CPF, secret)
	if !b {
		log.Fatal("Erro na autenticação")
	}
	os.Setenv("ACCESS_SECRET", "asdhjkasjheee")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": CPF,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		log.Fatal("Erro ao gerar token")
		return "", err
	}

	return tokenString, nil
}

func Authorize(token *jwt.Token) interface{} {
	fmt.Println("Token no authorize: ", token)
	
	var accessUuID string
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		accessUuID, _ = claims["user"].(string)

	}
	return accessUuID

}
