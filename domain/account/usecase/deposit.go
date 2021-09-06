package usecase

import (
	"errors"
	"fmt"
)

func (c AccountUc) Deposit(CPF string, amount float64) error {

	account, err := c.r.FindOne(CPF)
	fmt.Printf("CPF: %s", CPF)
	if err != nil {
		fmt.Println("the account does not exist")
		//TODO add a new sentinel
		return errors.New("the account does not exist")
	}
	account.Balance += amount
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("conta de destino a ser atualizada no UC %s", account.ID)
	return nil

}
