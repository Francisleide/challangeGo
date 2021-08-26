package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	middleware "github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
)

type DepositInput struct {
	Amount float64 `json: "amount"`
}

func ToDeposit(serv *mux.Router, usecase account.UseCase) *Handler {
	h := &Handler{
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
func (h Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	var deposit DepositInput
	usr, ok := middleware.GetAccountID(r.Context())
	if !ok {
		log.Fatal("unauthenticated user")
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deposit)
	if err != nil {
		log.Fatal("error ", err)
	}
	h.account.Deposit(usr, deposit.Amount)

	if err != nil {
		log.Fatal(err)
	}

}
