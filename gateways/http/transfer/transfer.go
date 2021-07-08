package transfer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/entities"
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

func (h Handler) Create_transfer(w http.ResponseWriter, r *http.Request) {
	var tr entities.TransferInput
	accountId, ok := middlware.GetAccountID(r.Context())
	if !ok || accountId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.Response.StatusCode)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&tr)
	fmt.Println("Na rota, ammount: ", tr.Ammount)
	fmt.Println("Id na rota: ", accountId)
	if err != nil {
		log.Fatal("Erro na hora de pegar elementos do body: ", err)
	}
	_, erro := h.transfer.Create_transfer(accountId, tr.Cpf_destino, tr.Ammount)
	if erro != nil {
		fmt.Println("Erro! Saldo insuficiente")
		w.WriteHeader(r.Response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
