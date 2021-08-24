package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/account"
	"github.com/francisleide/ChallangeGo/gateways/http/middlware"
	"github.com/gorilla/mux"
)

type HandlerWithdraw struct {
	account account.UseCase
}

type Withdraw struct {
	Amount float64 `json: "amount"`
}

func ToWithdraw(serv *mux.Router, usecase account.UseCase) *HandlerWithdraw {
	h := &HandlerWithdraw{
		account: usecase,
	}
	serv.HandleFunc("/withdraw", h.Withdraw).Methods("Post")
	return h
}

// ShowAccount godoc
// @Summary Make a Withdraw
// @Description Make a Withdraw from an authentic account
// @Param Body body Withdraw true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "x-request-id"
// @Router /withdraw [post]
func (h HandlerWithdraw) Withdraw(w http.ResponseWriter, r *http.Request) {
	var withdraw Withdraw

	accountID, _ := middlware.GetAccountID(r.Context())

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&withdraw)
	if err != nil {
		log.Fatal(err)
	}
	//mudar para account
	ok := h.account.WithDraw(accountID, withdraw.Amount)
	if !ok {
		log.Panic("Saldo insuficiente!")
	}

}
