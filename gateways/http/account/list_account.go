package account

import (
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	ID        string
	Name      string
	CPF       string
	Balance   float64
	CreatedAt string
}

// ShowAccount godoc
// @Summary Get accounts
// @Description List all accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} []Account
// @Failure 400 "Failed to decode"
// @Failure 500 "Unexpected internal server error"
// @Router /accounts [GET]
func (h Handler) ListAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := h.account.ListAll()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		log.Fatal(err)
	}

}
