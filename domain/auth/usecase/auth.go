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
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Claims struct {
	Cpf string `json:"cpf"`
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

func (a AuthUc) Login(cpf, secret string) bool {
	var account entities.Account
	acc := a.r.FindOne(account.Cpf)
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

func (a AuthUc) CreateToken(cpf string, secret string) (string, error) {
	//var err error
	//Creating Access Token
	b := a.Login(cpf, secret)
	if !b {
		log.Fatal("Erro na autenticação")
	}
	os.Setenv("ACCESS_SECRET", "asdhjkasjheee")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": cpf,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		log.Fatal("Erro ao gerar token")
		return "", err
	}

	return tokenString, nil
	/*os.Setenv("ACCESS_SECRET", "cfdfgdfg") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["Id"] = cpf
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil*/

}

func Authorize(token *jwt.Token) interface{} {
	fmt.Println("Token no authorize: ", token)
	
	var accessUuid string
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		fmt.Println("claims ok!")
		accessUuid, _ = claims["user"].(string)

	}
	fmt.Println("O id: ", accessUuid)
	return accessUuid

}
