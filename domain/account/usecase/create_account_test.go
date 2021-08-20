package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallangeGo/domain/account/usecase"
	"github.com/francisleide/ChallangeGo/domain/entities"
)

type repoMock struct {
	Account entities.Account
}

func (r *repoMock) ListAllAccounts() []entities.Account {
	return nil
}

func (r *repoMock) FindOne(CPF string) entities.Account {
	return entities.Account{
		Balance: 100.0,
		CPF:     "12345679210",
	}
}
func (r *repoMock) UpdateBalance(account entities.Account) {
	r.Account = account
}
func (r *repoMock) InsertAccount(accountInput entities.AccountInput) (*entities.Account, error) {
	if accountInput.CPF == "12345678911" {
		return nil, errors.New("Conta já existente")
	}
	return &r.Account, nil
}

func TestCreateAccount(t *testing.T) {

	var account entities.Account
	account.CPF = "12345678910"
	t.Run("Cpf não existe no banco e a conta é criada", func(t *testing.T) {
		r := repoMock{}
		r.Account = account
		a := usecase.NewAccountUc(&r)
		var accountInput entities.AccountInput
		accountInput.CPF = "12345678910"
		accountInput.Nome = "Teste"
		accountInput.Secret = "111"
		accountRecived, err := a.CreateAccount(accountInput)
		if err != nil {
			t.Errorf("Erros: %s", err)
		}
		if accountRecived.CPF != "12345678910" {
			t.Errorf("Esperado %s e recebido %s", accountInput.CPF, accountRecived.CPF)
		}

	})

	t.Run("cpf existe no banco e a conta não é criada", func(t *testing.T) {
		r := repoMock{}
		a := usecase.NewAccountUc(&r)
		var accountInput entities.AccountInput
		accountInput.CPF = "12345678911"
		accountInput.Nome = "Teste"
		accountInput.Secret = "111"
		_, err := a.CreateAccount(accountInput)
		if err == nil {
			t.Error("Já existia uma conta com o mesmo cpf e a conta foi criada", err)
		}

	})
}
