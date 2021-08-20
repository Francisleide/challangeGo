package usecase_test

import (
	"testing"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/domain/transfer/usecase"
)

type repoMock struct {
	Transfer entities.Transfer
	Account entities.Account
}

func (r *repoMock) InsertTransfer(account_origem, account_destino entities.Account, amount float64) (*entities.Transfer, error) {
	newAccount := &r.Account
	newAccount.Balance -= amount
	return &r.Transfer, nil
}
func (r *repoMock) FindOne(cpf string) entities.Account {
	var account entities.Account
	account.Balance = 100
	account.CPF = "231"
	return account
}

func Test_Create_transfer(t *testing.T) {
	//Caso 1: O emissário tem saldo na conta e a transferência ocorre
	//Caso 2: O emissário não tem saldo na conta e a transferência não ocorre
	var account entities.Account
	account.Balance = 100
	account.CPF = "231"
	
	var tt entities.Transfer
	tt.AccountDestinationID = "123"
	tt.AccountOriginID = account.CPF
	tt.Amount = 10.0
	t.Run("O emissário tem saldo na conta e a transferência ocorre", func(t *testing.T) {
		r := repoMock{
			Transfer: tt,
			Account:  account,
		}

		transfer:=usecase.NewTransfer(&r)
		_, err:=transfer.CreateTransfer(tt.AccountOriginID, tt.AccountDestinationID, tt.Amount)
		balanceEsperado:= 90.0
		if r.Account.Balance != balanceEsperado{
			t.Errorf("Balance esperado na conta : %.2f, Balance recebido: %.2f", balanceEsperado, r.Account.Balance)
		}
		if err != nil{
			t.Errorf("A operação deveria ocorrer sem erros, mas foi capturado: %s", err)
		}
	})
}
