package transfer

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	"github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	transfer transfer.UseCase
	account  account.UseCase
	log      *logrus.Entry
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

func NewTransfer(serv *mux.Router, usecase transfer.UseCase, accountUC account.UseCase, log *logrus.Entry) *Handler {
	h := &Handler{
		transfer: usecase,
		account:  accountUC,
		log:      log,
	}

	serv.HandleFunc("/transfers", h.CreateTransfer).Methods("Post")
	serv.HandleFunc("/transfers", h.ListUserTransfers).Methods("Get")

	return h
}

// CreateTransfer godoc
// @Summary Make a transfer
// @Description Transfer between accounts. The account that will make the transfer must be authenticated with a token.
// @Param Authorization header string true "Bearer"
// @Accept  json
// @Produce  json
// @Success 201 "Created"
// @Failure 401 "Invalid or missing token"
// @Failure 404 "Account origin/destination not found"
// @Router /transfers [post]
func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var tr TransferInput
	CPF, ok := middleware.GetCPF(r.Context())

	if !ok {
		h.log.Errorln("incorrect token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if CPF == "" {
		h.log.Errorln("account not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tr)

	if err != nil {
		h.log.WithError(err).Error("problems finding the source cpf")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	accountOrigin, errOrigin := h.account.GetAccountByCPF(CPF)
	if errOrigin != nil {
		h.log.WithError(errOrigin).Error("problems finding the source cpf")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	accountDestination, errorDestine := h.account.GetAccountByID(tr.AccountDestinationID)
	if errorDestine != nil {
		h.log.WithError(errorDestine).Error("problems finding the destination cpf")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_,err = h.account.Withdraw(CPF, tr.Amount)
	if err != nil {
		h.log.WithError(err).Error("error when making the withdrawal")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_,err = h.account.Deposit(accountDestination.CPF, tr.Amount)
	if err != nil {
		h.log.WithError(err).Error("error when making the deposit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	transfer, errCreateTransfer := h.transfer.CreateTransfer(accountOrigin, accountDestination, tr.Amount)
	if errCreateTransfer != nil {
		h.log.WithError(errCreateTransfer).Error("error creating transfer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transfer)
	h.log.Info("transfer successful")

}

// ListUserTransfers godoc
// @Summary List transfers from a user
// @Description Lists all transfers made by an authenticated user
// @Param Authorization header string true "Bearer"
// @Accept  json
// @Produce  json
// @Success 200 {object} []Transfer
// @Failure 401 "Invalid or missing token"
// @Failure 500 "Problems finding the transfers"
// @Router /transfers [get]
func (h Handler) ListUserTransfers(w http.ResponseWriter, r *http.Request) {
	accountCPF, ok := middleware.GetCPF(r.Context())
	if !ok {
		h.log.Errorln("incorrect token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if accountCPF == "" {
		h.log.Errorln("account not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	account, err := h.account.GetAccountByCPF(accountCPF)
	if err != nil {
		h.log.WithError(err).Error("problems finding the source cpf")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	transfers, errList := h.transfer.ListUserTransfers(account.ID)
	if errList != nil {
		h.log.WithError(err).Error("problems finding the transfers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errList = json.NewEncoder(w).Encode(transfers)
	if errList != nil {
		h.log.WithError(errList).Errorln("failed to perform transfer encode")
		w.WriteHeader(http.StatusBadRequest)
	}
	h.log.Info("transfers listed successfully")

}
