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

type JSONFormatter struct {
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
// @Success 200
// @Router /transfers [post]
func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var tr TransferInput
	CPF, ok := middleware.GetCPF(r.Context())

	if !ok {
		h.log.Errorln("incorrect token")
		w.Header().Set("Content-Type", "application/json")
		return
	}
	if CPF == "" {
		h.log.Errorln("account not found")
		w.Header().Set("Content-Type", "application/json")
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tr)

	if err != nil {
		h.log.WithError(err).Error("problems finding the source cpf")
		return
	}
	accountOrigin, errOrigin := h.account.GetAccountByCPF(CPF)
	if errOrigin != nil {
		h.log.WithError(errOrigin).Error("problems finding the source cpf")
		return
	}

	accountDestination, errorDestine := h.account.GetAccountByID(tr.AccountDestinationID)
	if errorDestine != nil {
		h.log.WithError(errorDestine).Error("problems finding the destination cpf")
		return
	}
	err = h.account.Withdraw(CPF, tr.Amount)
	if err != nil {
		h.log.WithError(err).Error("error when making the withdrawal")
		return
	}

	err = h.account.Deposit(accountDestination.CPF, tr.Amount)
	if err != nil {
		h.log.WithError(err).Error("error when making the deposit")
		return
	}

	_, errCreateTransfer := h.transfer.CreateTransfer(accountOrigin, accountDestination, tr.Amount)
	if errCreateTransfer != nil {
		h.log.WithError(errCreateTransfer).Error("error creating transfer")
		return
	}
	h.log.Info("transfer successful")

	w.Header().Set("Content-Type", "application/json")
}

// ListUserTransfers godoc
// @Summary List transfers from a user
// @Description Lists all transfers made by an authenticated user
// @Param Authorization header string true "Bearer"
// @Accept  json
// @Produce  json
// @Success 200 {object} []Transfer
// @Router /transfers [get]
func (h Handler) ListUserTransfers(w http.ResponseWriter, r *http.Request) {
	accountCPF, ok := middleware.GetCPF(r.Context())
	if !ok {
		h.log.Errorln("incorrect token")
		return
	}
	if accountCPF == "" {
		h.log.Errorln("account not found")
		return
	}
	account, err := h.account.GetAccountByCPF(accountCPF)
	if err != nil {
		h.log.WithError(err).Error("problems finding the source cpf")
		return
	}

	transfers, errList := h.transfer.ListUserTransfers(account.ID)
	if errList != nil {
		h.log.WithError(err).Error("problems finding the transfers")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	errList = json.NewEncoder(w).Encode(transfers)
	if errList != nil {
		h.log.WithError(errList).Errorln("failed to perform transfer encode")
		w.WriteHeader(http.StatusBadRequest)
	}
	h.log.Info("transfers listed successfully")

}
