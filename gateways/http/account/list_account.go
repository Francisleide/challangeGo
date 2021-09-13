package account

import (
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	ID        string  `json: "id"`
	Name      string  `json: "name"`
	CPF       string  `json: "cpf"`
	Balance   float64 `json: "balance"`
	CreatedAt string  `json: "createdat"`
}

// ListAllAcounts godoc
// @Summary Get accounts
// @Description List all accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} []Account
// @Router /accounts [GET]
func (h Handler) ListAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := h.account.ListAll()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		log.Fatal(err)
	}

}
