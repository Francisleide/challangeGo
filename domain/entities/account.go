package entities

import (
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/satori/uuid.go"
	"golang.org/x/crypto/bcrypt"
)

type TransactionOutput struct {
	ID              string
	PreviousBalance float64
	ActualBalance   float64
}
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

func EncryptSecret(pass string) (string, error) {
	secret, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(secret), nil

}

func ValidateCPF(CPF string) bool {
	if len(CPF) == 11 {
		runes := []rune(CPF)
		firstBuf := make([]byte, 1)
		_ = utf8.EncodeRune(firstBuf, runes[9])
		firstDigit, _ := strconv.Atoi(string(firstBuf))
		secondBuf := make([]byte, 1)
		_ = utf8.EncodeRune(secondBuf, runes[10])
		secondDigit, _ := strconv.Atoi(string(secondBuf))
		totalFirst := 0
		value := 10
		for i := 0; i < 9; i++ {
			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, runes[i])
			digit, _ := strconv.Atoi(string(buf))
			totalFirst += value * digit
			value--
		}
		result := (totalFirst * 10) % 11
		if result == 10 {
			result = 0
		}
		if result == int(firstDigit) {
			totalSecond := 0
			value = 11
			for i := 0; i < 10; i++ {
				buf := make([]byte, 1)
				_ = utf8.EncodeRune(buf, runes[i])
				digit2, _ := strconv.Atoi(string(buf))
				totalSecond += value * digit2
				value--
			}
			result2 := (totalSecond * 10) % 11
			if result2 == 10 {
				result2 = 0
			}
			if result2 == secondDigit {
				return true
			}
		}

	}
	return false
}

func ValidateSecret(secret string) bool {
	numbers := 0
	if len(secret) > 4 {
		for _, char := range secret {
			for _, number := range "1234567890" {
				if char == number {
					numbers++
				}
			}
		}
	}
	if numbers > 0 && numbers < len(secret) {
		return true
	}
	return false
}

func ValidateName(name string) bool {
	return len(name) != 0
}

func NewAccount(name, CPF, secret string) (Account, error) {
	if !ValidateSecret(secret) {
		return Account{}, ErrorInvalidSecret
	}

	if !ValidateName(name) {
		return Account{}, ErrorNameEmpty
	}
	if !ValidateCPF(CPF) {
		return Account{}, ErrorInvalidCPF
	}

	secret, err := EncryptSecret(secret)
	if err != nil {
		return Account{}, nil
	}
	return Account{
		ID:        uuid.NewV4().String(),
		Name:      name,
		CPF:       CPF,
		Secret:    secret,
		CreatedAt: time.Now().UTC().String(),
	}, nil
}
