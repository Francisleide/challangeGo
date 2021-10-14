package usecase

func (c AccountUc) GetBalance(accountID string) (float64, error) {
	account, err := c.r.FindByID(accountID)
	if err != nil {
		return 0, ErrorRetrieveAccount
	}
	return account.Balance, nil
}
