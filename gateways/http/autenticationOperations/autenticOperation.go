package autenticationoperations

import (
	autenticoperations "github.com/francisleide/ChallangeGo/domain/autenticOperations"
	"github.com/gorilla/mux"
)

type Handler struct {
	autentic autenticoperations.UseCase
}

func AutenticationOperations(serv *mux.Router, usecase autenticoperations.UseCase) *Handler {
	h := &Handler{
		autentic: usecase,
	}
	serv.HandleFunc("/withdraw", h.Withdraw).Methods("Post")
	serv.HandleFunc("/deposite", h.Deposite).Methods("Post")

	return h
}
