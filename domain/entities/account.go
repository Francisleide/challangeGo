package entities

import (
	"log"
	"time"

	uuid "github.com/satori/uuid.go"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id         string  `json: "account_id"`
	Nome       string  `json: "nome"`
	Cpf        string  `json: "cpf"`
	Secret     string  `json: "secret"`
	Balance    float64 `json: "balance"`
	Created_at string  `json: "created_at"`
}

type AccountInput struct {
	Nome   string `json: "nome"`
	Cpf    string `json: "cpf"`
	Secret string `json: "secret"`
}

func GenerateId() string {
	return uuid.NewV4().String()
}

func EncryptSecret(pass string) (string, error) {
	secret, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(secret), nil

}
func NewAccount(nome, cpf, secret string) Account {
	secret, err := EncryptSecret(secret)
	if err != nil {
		log.Fatal(err)
	}
	return Account{
		Id:         GenerateId(),
		Nome:       nome,
		Cpf:        cpf,
		Secret:     secret,
		Created_at: time.Now().Format(time.RFC822),
	}
}
