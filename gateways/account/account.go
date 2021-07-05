package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/account"
	"github.com/francisleide/ChallangeGo/entities"
	"github.com/gorilla/mux"
)

type Handler struct {
	account account.UseCase
}

func Accounts(serv *mux.Router, usecase account.UseCase) *Handler {
	h := &Handler{
		account: usecase,
	}

	serv.HandleFunc("/accounts", h.Create_account).Methods("Post")
	serv.HandleFunc("/accounts", h.List_all_accounts).Methods("Get")
	//serv.HandleFunc("/withdraw", h.Withdraw).Methods("Post")


	return h
}

func (h Handler) Create_account(w http.ResponseWriter, r *http.Request) {
	var acc entities.AccountInput
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&acc)
	fmt.Println("CPF na rota: ", acc.Cpf)
	if err != nil {
		//implenmentar erro aqui
	}
	h.account.Create_account(entities.AccountInput{
		Nome:   acc.Nome,
		Cpf:    acc.Cpf,
		Secret: acc.Secret,
	})
	w.Header().Set("Content-Type", "application/json")

	//return json.NewEncoder(w).Encode(a)
}
