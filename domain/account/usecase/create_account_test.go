package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallangeGo/domain/account/usecase"
	"github.com/francisleide/ChallangeGo/domain/entities"
)

type repo_mock struct {
	Account entities.Account
}

func (r *repo_mock) List_all_accounts() []entities.Account{
	return nil
}

func (r *repo_mock) FindOne(cpf string) entities.Account {
	return entities.Account{
		Balance: 100.0,
		Cpf:     "12345679210",
	}
}
func (r *repo_mock) UpdateBalance(account entities.Account) {
	r.Account = account
}
func (r *repo_mock) InsertAccount(accountInput entities.AccountInput) (*entities.Account, error) {
	if accountInput.Cpf == "12345678911"{
		return nil, errors.New("Conta já existente")
	} 
	return &r.Account, nil
}

func Test_Create_account(t *testing.T) {
	//Caso1: cpf não existe no banco e a conta é criada
	//Caso2: cpf existe no banco e a conta não é criada

	var account entities.Account
	account.Cpf = "12345678910"
	t.Run("Cpf não existe no banco e a conta é criada", func(t *testing.T) {
		r := repo_mock{}
		r.Account = account
		a := usecase.NewAccountUc(&r)
		var accountInput entities.AccountInput
		accountInput.Cpf = "12345678910"
		accountInput.Nome = "Teste"
		accountInput.Secret = "111"
		account_recebido, err := a.Create_account(accountInput)
		if err != nil {
			t.Errorf("Erros: %s", err)
		}
		if account_recebido.Cpf != "12345678910" {
			t.Errorf("Esperado %s e recebido %s", accountInput.Cpf, account_recebido.Cpf)
		}

	})

	t.Run("cpf existe no banco e a conta não é criada", func(t *testing.T) {
		r := repo_mock{}
		a := usecase.NewAccountUc(&r)
		var accountInput entities.AccountInput
		accountInput.Cpf = "12345678911"
		accountInput.Nome = "Teste"
		accountInput.Secret = "111"
		_, err := a.Create_account(accountInput)
		if err == nil{
			t.Error("Já existia uma conta com o mesmo cpf e a conta foi criada", err)
		}
		
	})
}
