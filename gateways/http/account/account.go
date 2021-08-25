package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/gorilla/mux"
)

type Handler struct {
	account account.UseCase
}

func Accounts(serv *mux.Router, usecase account.UseCase) *Handler {
	h := &Handler{
		account: usecase,
	}

	serv.HandleFunc("/accounts", h.CreateAccount).Methods("Post")
	serv.HandleFunc("/accounts", h.ListAllAccounts).Methods("Get")

	return h
}

// ShowAccount godoc
// @Summary Create an account
// @Description Create an account with the basic information
// @Param Body body entities.AccountInput true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "request-id"
// @Router /accounts [post]
func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc entities.AccountInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&acc)
	if err != nil {
		//implenmentar erro aqui
	}
	h.account.CreateAccount(entities.AccountInput{
		Name:   acc.Name,
		CPF:    acc.CPF,
		Secret: acc.Secret,
	})
	w.Header().Set("Content-Type", "application/json")

}
