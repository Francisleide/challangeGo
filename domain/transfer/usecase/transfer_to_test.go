package usecase_test

import (
	"testing"

	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer/usecase"
)

type repoMock struct {
	Transfer entities.Transfer
	Account  entities.Account
}

func (r *repoMock) InsertTransfer(originAccount, destineAccount entities.Account, amount float64) (*entities.Transfer, error) {
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
	var account entities.Account
	account.Balance = 100
	account.CPF = "231"

	var tt entities.Transfer
	tt.DestineAccountID = "123"
	tt.OriginAccountID = account.CPF
	tt.Amount = 10.0
	t.Run("The emissary has a balance in the account and the transfer runs.", func(t *testing.T) {
		r := repoMock{
			Transfer: tt,
			Account:  account,
		}

		transfer := usecase.NewTransferUC(&r)
		_, err := transfer.CreateTransfer(tt.OriginAccountID, tt.DestineAccountID, tt.Amount)
		expectedBalance := 90.0
		if r.Account.Balance != expectedBalance {
			t.Errorf("Expected balance: %.2f, recived balance: %.2f", expectedBalance, r.Account.Balance)
		}
		if err != nil {
			t.Errorf("The operation should go without error, but it was captured: %s", err)
		}
	})
}
