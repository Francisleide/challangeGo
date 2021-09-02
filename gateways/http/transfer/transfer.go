package transfer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	"github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
)

type Handler struct {
	transfer transfer.UseCase
}

func Transfer(serv *mux.Router, usecase transfer.UseCase) *Handler {
	h := &Handler{
		transfer: usecase,
	}

	serv.HandleFunc("/transfers", h.CreateTransfer).Methods("Post")
	serv.HandleFunc("/transfers", h.ListUserTransfers).Methods("Get")

	return h
}

// ShowAccount godoc
// @Summary Make a transfer
// @Description Transfer between accounts. The account that will make the transfer must be authenticated with a token.
// @Param Body body entities.TransferInput true "Body"
// @Param Authorization header string true "Bearer Authorization Token"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "request-id"
// @Router /transfer [post]
func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var tr entities.TransferInput
	accountID, ok := middleware.GetAccountID(r.Context())
	if !ok || accountID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tr)
	if err != nil {
		log.Fatal(err)
	}
	_, error := h.transfer.CreateTransfer(accountID, tr.AccountDestinationID, tr.Amount)
	if error != nil {
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// ShowAccount godoc
// @Summary Make a transfer
// @Description Transfer between accounts. The account that will make the transfer must be authenticated with a token.
// @Param Body body []entities.Transfer true "Body"
// @Param Authorization header string true "Bearer Authorization Token"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "request-id"
// @Router /transfer [post]
func (h Handler) ListUserTransfers(w http.ResponseWriter, r *http.Request) {
	accountCPF, ok := middleware.GetAccountID(r.Context())
	if !ok || accountCPF == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transfers, err := h.transfer.ListUserTransfers(accountCPF)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transfers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
