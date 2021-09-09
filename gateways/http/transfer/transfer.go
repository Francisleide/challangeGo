package transfer

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	"github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
)

type Handler struct {
	transfer transfer.UseCase
	account  account.UseCase
}

type TransferInput struct {
	AccountDestinationID string  `json: "accountdestinationid"`
	Amount               float64 `json: "amount"`
}
type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               float64
	CreatedAt            string
}

func NewTransfer(serv *mux.Router, usecase transfer.UseCase, accountUC account.UseCase) *Handler {
	h := &Handler{
		transfer: usecase,
		account:  accountUC,
	}

	serv.HandleFunc("/transfers", h.CreateTransfer).Methods("Post")
	serv.HandleFunc("/transfers", h.ListUserTransfers).Methods("Get")

	return h
}

// ShowAccount godoc
// @Summary Make a transfer
// @Description Transfer between accounts. The account that will make the transfer must be authenticated with a token.
// @Param Authorization header string true "Bearer Authorization Token"
// @Accept  json
// @Produce  json
// @Success 200 {object} Transfer
// @Failure 400 "Failed to decode"
// @Failure 401 "Unauthorized"
// @Failure 500 "Unexpected internal server error"
// @Router /transfers [post]
func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var tr TransferInput
	accountID, ok := middleware.GetAccountID(r.Context())
	if !ok || accountID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tr)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}
	accountOrigin, errOrigin := h.account.GetAccountByCPF(accountID)
	if errOrigin != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	accountDestination, errorDestine := h.account.GetAccountByID(tr.AccountDestinationID)
	if errorDestine != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}
	err = h.account.Withdraw(accountID, tr.Amount)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	err = h.account.Deposit(accountDestination.CPF, tr.Amount)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	_, error := h.transfer.CreateTransfer(accountOrigin, accountDestination, tr.Amount)
	if error != nil {
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// ShowAccount godoc
// @Summary List transfers from a user
// @Description Lists all transfers made by an authenticated user
// @Param Authorization header string true "Bearer Authorization Token"
// @Accept  json
// @Produce  json
// @Success 200 {object} []Transfer
// @Failure 400 "Failed to decode"
// @Failure 401 "Unauthorized"
// @Failure 500 "Unexpected internal server error"
// @Router /transfers [get]
func (h Handler) ListUserTransfers(w http.ResponseWriter, r *http.Request) {
	accountCPF, ok := middleware.GetAccountID(r.Context())
	if !ok || accountCPF == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	account, err := h.account.GetAccountByCPF(accountCPF)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transfers, errList := h.transfer.ListUserTransfers(account.ID)
	if errList != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	errList = json.NewEncoder(w).Encode(transfers)
	if errList != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
