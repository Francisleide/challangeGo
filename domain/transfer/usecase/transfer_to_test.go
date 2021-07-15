package usecase_test

import (
	"testing"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/domain/transfer/usecase"
)

type repo_mock struct {
	Transfer entities.Transfer
	Account entities.Account
}

func (r *repo_mock) InsertTransfer(account_origem, account_destino entities.Account, amount float64) (*entities.Transfer, error) {
	newAccount := &r.Account
	newAccount.Balance -= amount
	return &r.Transfer, nil
}
func (r *repo_mock) FindOne(cpf string) entities.Account {
	var account entities.Account
	account.Balance = 100
	account.Cpf = "231"
	return account
}

func Test_Create_transfer(t *testing.T) {
	//Caso 1: O emissário tem saldo na conta e a transferência ocorre
	//Caso 2: O emissário não tem saldo na conta e a transferência não ocorre
	var account entities.Account
	account.Balance = 100
	account.Cpf = "231"
	
	var tt entities.Transfer
	tt.Account_destination_id = "123"
	tt.Account_origin_id = account.Cpf
	tt.Amount = 10.0
	t.Run("O emissário tem saldo na conta e a transferência ocorre", func(t *testing.T) {
		r := repo_mock{
			Transfer: tt,
			Account:  account,
		}

		transfer:=usecase.NewTransfer(&r)
		_, err:=transfer.Create_transfer(tt.Account_origin_id, tt.Account_destination_id, tt.Amount)
		balanceEsperado:= 90.0
		if r.Account.Balance != balanceEsperado{
			t.Errorf("Balance esperado na conta : %.2f, Balance recebido: %.2f", balanceEsperado, r.Account.Balance)
		}
		if err != nil{
			t.Errorf("A operação deveria ocorrer sem erros, mas foi capturado: %s", err)
		}
	})
}
