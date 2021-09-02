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

func (r *repoMock) ListUserTransfers(accountID string) ([]entities.Transfer, error) {
	return []entities.Transfer{}, nil
}
func (r *repoMock) FindByID(accountID string) (entities.Account, bool) {
	return entities.Account{}, true
}
func (r *repoMock) InsertTransfer(transfer entities.Transfer) (entities.Transfer, error) {
	newAccount := &r.Account
	newAccount.Balance -= transfer.Amount
	return r.Transfer, nil
}
func (r *repoMock) FindOne(cpf string) (entities.Account, bool) {
	var account entities.Account
	account.Balance = 100
	account.CPF = "231"
	return account, true
}
func (r *repoMock) UpdateBalance(account entities.Account) bool {
	return true
}

func Test_Create_transfer(t *testing.T) {
	var account entities.Account
	account.Balance = 100
	account.CPF = "231"

	var tt entities.Transfer
	tt.AccountDestinationID = "123"
	tt.AccountOriginID = account.CPF
	tt.Amount = 10.0
	t.Run("The emissary has a balance in the account and the transfer runs.", func(t *testing.T) {
		r := repoMock{
			Transfer: tt,
			Account:  account,
		}

		transfer := usecase.NewTransferUC(&r)
		_, err := transfer.CreateTransfer(tt.AccountOriginID, tt.AccountDestinationID, tt.Amount)
		expectedBalance := 90.0
		if r.Account.Balance != expectedBalance {
			t.Errorf("Expected balance: %.2f, received balance: %.2f", expectedBalance, r.Account.Balance)
		}
		if err != nil {
			t.Errorf("The operation should go without error, but it was captured: %s", err)
		}
	})
}
