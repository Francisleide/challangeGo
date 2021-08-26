package entities

import (
	"log"
	"time"

	"github.com/satori/uuid.go"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt string
}

type AccountInput struct {
	Name   string
	CPF    string
	Secret string
}

func GenerateID() string {
	return uuid.NewV4().String()
}

func EncryptSecret(pass string) (string, error) {
	secret, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(secret), nil

}
func NewAccount(name, CPF, secret string) Account {
	secret, err := EncryptSecret(secret)
	if err != nil {
		log.Fatal(err)
	}
	return Account{
		ID:        GenerateID(),
		Name:      name,
		CPF:       CPF,
		Secret:    secret,
		CreatedAt: time.Now().Format(time.RFC822),
	}
}
