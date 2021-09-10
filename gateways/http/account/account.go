package account

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/gorilla/mux"
)

type Handler struct {
	account account.UseCase
}
type AccountInput struct {
	Name   string `json: "name"`
	CPF    string `json: "cpf"`
	Secret string `json: "secret"`
}

func Accounts(serv *mux.Router, usecase account.UseCase) *Handler {
	h := &Handler{
		account: usecase,
	}

	serv.HandleFunc("/accounts", h.CreateAccount).Methods("Post")
	serv.HandleFunc("/accounts", h.ListAllAccounts).Methods("Get")
	serv.HandleFunc("/accounts/{id}/balance", h.GetBalance).Methods(("Get"))

	return h
}

// ShowAccount godoc
// @Summary Create an account
// @Description Create an account with the basic information
// @Param Body body AccountInput true "Body"
// @Accept  json
// @Produce  json
// @Success 200 {object} Account
// @Router /accounts [post]
func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountInput AccountInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	account, err := h.account.CreateAccount(entities.AccountInput{
		Name:   accountInput.Name,
		CPF:    accountInput.CPF,
		Secret: accountInput.Secret,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
