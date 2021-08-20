package usecase_test

import (
	"testing"

	"github.com/francisleide/ChallangeGo/domain/autenticOperations/usecase"
	"github.com/francisleide/ChallangeGo/domain/entities"
)


type repoMock struct {
	Account entities.Account
}

func (r *repoMock) ListAllAccounts() []entities.Account{
	return nil
}
func (r *repoMock) FindOne(cpf string) entities.Account {
	return entities.Account{
		Balance: 100.0,
		CPF:     "12345679210",
	}
}
func (r *repoMock) UpdateBalance(account entities.Account) {
	r.Account = account
}
func (r *repoMock) InsertAccount(aaccountInput entities.AccountInput) (*entities.Account, error) {
	return &entities.Account{}, nil
}

func Test_WithDraw(t *testing.T) {
	//Caso 1: Tem saldo suficiente na conta e consegue sacar
	//Caso 2: Não tem saldo suficiente e não consegue sacar
	var account entities.Account
	account.Balance = 100
	account.CPF = "12345679210"

	t.Run("Tem saldo suficiente na conta e consegue sacar", func(t *testing.T) {
		amount := 50.0
		r := repoMock{}
		r.Account = account
		autentic := usecase.NewAutentic(&r)

		ok := autentic.WithDraw(r.Account.CPF, amount)

		esperado := true
		if esperado != ok {
			t.Error("O saque não teve sucesso!")
		}
		balanceEsperado := 50.0
		if balanceEsperado != r.Account.Balance {
			t.Errorf("O saldo não foi atualizado com o valor esperado. Era esperado %.2f e recebido %.2f ", balanceEsperado, r.Account.Balance)
		}

	})
	t.Run("Não tem saldo suficiente e não consegue sacar", func(t *testing.T) {
		amount := 120.0

		r := repoMock{}
		r.Account = account
		autentic := usecase.NewAutentic(&r)

		ok := autentic.WithDraw(r.Account.CPF, amount)

		esperado := false
		if esperado != ok {
			t.Error("Era esperado que o saque não fosse efetuado e o saque ocorreu.")
		}
		balanceEsperado := 100.0
		if balanceEsperado != r.Account.Balance {
			t.Errorf("Era esperado %.2f e recebido %.2f ", balanceEsperado, r.Account.Balance)
		}

	})
}

func Test_Deposite(t *testing.T) {
	//Caso 1: Depósito de um valor e valor acrescido à conta
	//Caso 2: Depósito de um valor negativo e o valor da conta permanece o mesmo
	var account entities.Account
	account.Balance = 100
	account.CPF = "12345679210"

	t.Run("Depósito de um valor e valor acrescido à conta", func(t *testing.T) {
		r := repoMock{}
		r.Account = account
		a:=usecase.NewAutentic(&r)
		amount:= 100.0

		a.Deposite(account.CPF, amount)
		balanceEsperado := 200.0
		if balanceEsperado != r.Account.Balance{
			t.Errorf("Balance esperado %.2f e Balance retornado %.2f", balanceEsperado, r.Account.Balance)
		}
		
	})

	t.Run("Depósito de um valor negativo e o valor da conta permanece o mesmo", func(t *testing.T) {
		r := repoMock{}
		a:= usecase.NewAutentic(&r)
		r.Account = account
		a.Deposite(account.CPF, -150.0)
		balanceEsperado := 100.0
		if balanceEsperado != r.Account.Balance{
			t.Errorf("Balance esperado %.2f e Balance retornado %.2f", balanceEsperado, r.Account.Balance)
		}

	})
}
