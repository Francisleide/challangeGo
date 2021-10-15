package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountBalance struct {
	Balance float64 `json:"balance"`
}

// GetBalance godoc
// @Summary account balance
// @Description Show the balance of a specific account
///@Accept  json
// @Produce  json
// @Success 200 {object} AccountBalance
// @Failure 400 "Failed to decode"
// @Failure 401 "Invalid or missing token"
// @Failure 404 "Account not found"
// @Failure 500 "Unexpected internal server error"
// @Router /accounts/{id}/balance [GET]
func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["id"]
	balance, err := h.account.GetBalance(accountID)
	if err != nil {
		h.log.WithError(err).Errorf("failed to find balance")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var accountBalance AccountBalance
	accountBalance.Balance = balance
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(accountBalance)
	if err != nil {
		h.log.WithError(err).Errorf("failed to write json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
