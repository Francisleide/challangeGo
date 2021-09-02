package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account/usecase"
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type repoMock struct {
	Account entities.Account
}

func (r *repoMock) ListAllAccounts() []entities.Account {
	return nil
}

func (r *repoMock) FindOne(CPF string) (entities.Account, bool) {
	return entities.Account{
		Balance: 100.0,
		CPF:     "12345679210",
	}, true
}

func (r *repoMock) FindByID(accountID string) (entities.Account, bool) {
	return entities.Account{}, false
}
func (r *repoMock) UpdateBalance(account entities.Account) bool {
	r.Account = account
	return true
}
func (r *repoMock) InsertAccount(account entities.Account) error {
	if account.CPF == "12345678911" {
		return errors.New("the account already exists")
	}
	return nil
}

func TestCreateAccount(t *testing.T) {

	var account entities.Account
	account.CPF = "12345678910"
	t.Run("CPF does not exist in the bank and the account is created", func(t *testing.T) {
		r := repoMock{}
		r.Account = account
		a := usecase.NewAccountUc(&r)
		var accountInput entities.AccountInput
		accountInput.CPF = "12345678910"
		accountInput.Name = "Test"
		accountInput.Secret = "111"
		accountReceived, err := a.CreateAccount(accountInput)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if accountReceived.CPF != "12345678910" {
			t.Errorf("Expected %s and was received  %s", accountInput.CPF, accountReceived.CPF)
		}

	})

	t.Run("CPF exists in the bank and the account is not created", func(t *testing.T) {
		r := repoMock{}
		a := usecase.NewAccountUc(&r)
		var accountInput entities.AccountInput
		accountInput.CPF = "12345678911"
		accountInput.Name = "Test"
		accountInput.Secret = "111"
		_, err := a.CreateAccount(accountInput)
		if err == nil {
			t.Error("An account already existed with the same cpf and the account was created", err)
		}

	})
}
