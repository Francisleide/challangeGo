package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/account"
	"github.com/francisleide/ChallangeGo/domain/entities"
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

	return h
}

// ShowAccount godoc
// @Summary Create an account
// @Description Create an account with the basic information
// @Param Body body entities.AccountInput true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "x-request-id"
// @Router /accounts [post]
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

}
