package transfer

import (
	"encoding/json"
	"log"
	"net/http"

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

	serv.HandleFunc("/transfer", h.CreateTransfer).Methods("Post")

	return h
}

// ShowAccount godoc
// @Summary Make a transfer
// @Description Transfer between accounts. The account that will make the transfer must be authenticated with a token.
// @Param Body body TransferInput true "Body"
// @Param Authorization header string true "Bearer Authorization Token"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "request-id"
// @Router /transfer [post]
func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var tr transfer.TransferInput
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
	_, error := h.transfer.CreateTransfer(accountID, tr.DestinationCPF, tr.Amount)
	if error != nil {
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
