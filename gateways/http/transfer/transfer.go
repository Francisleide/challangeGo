package transfer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/transfer"
	"github.com/francisleide/ChallangeGo/gateways/http/middlware"
	"github.com/gorilla/mux"
)

type Handler struct {
	transfer transfer.UseCase
}

func Transfer(serv *mux.Router, usecase transfer.UseCase) *Handler {
	h := &Handler{
		transfer: usecase,
	}

	serv.HandleFunc("/transfer", h.Create_transfer).Methods("Post")

	return h
}

// ShowAccount godoc
// @Summary Make a transfer
// @Description Transfer between accounts. The account that will make the transfer must be authenticated with a token.
// @Param Body body TransferInput true "Body"
// @Param Authorization header string true "Bearer Authorization Token"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "x-request-id"
// @Router /transfer [post]
func (h Handler) Create_transfer(w http.ResponseWriter, r *http.Request) {
	var tr transfer.TransferInput
	accountId, ok := middlware.GetAccountID(r.Context())
	if !ok || accountId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&tr)
	fmt.Println("Na rota, ammount: ", tr.Amount)
	fmt.Println("Id na rota: ", accountId)
	if err != nil {
		log.Fatal("Erro na hora de pegar elementos do body: ", err)
	}
	_, erro := h.transfer.Create_transfer(accountId, tr.Cpf_destino, tr.Amount)
	if erro != nil {
		fmt.Println("Erro! Saldo insuficiente")
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
