package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	middleware "github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
)

type DepositHandler struct {
	account account.UseCase
}

type DepositInput struct {
	Amount float64 `json: "amount"`
}

func ToDeposit(serv *mux.Router, usecase account.UseCase) *DepositHandler {
	h := &DepositHandler{
		account: usecase,
	}
	serv.HandleFunc("/deposit", h.Deposit).Methods("Post")

	return h
}

// ShowAccount godoc
// @Summary Make a deposit
// @Description Make a deposit from an authentic account
// @Param Body body DepositInput true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "request-id"
// @Router /deposit [post]
func (h DepositHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var deposit DepositInput
	usr, ok := middleware.GetAccountID(r.Context())
	if !ok {
		log.Fatal("Usuário não autenticado: ")
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deposit)
	if err != nil {
		log.Fatal("Não consegui ler o body: ", err)
	}
	//mudar para account
	h.account.Deposit(usr, deposit.Amount)

	if err != nil {
		log.Fatal(err)
	}

}
