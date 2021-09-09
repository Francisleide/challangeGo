package account

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	middleware "github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
)

type Withdraw struct {
	Amount float64 `json: "amount"`
}

func ToWithdraw(serv *mux.Router, usecase account.UseCase) *Handler {
	h := &Handler{
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
// @Header 201 {string} Token "request-id"
// @Router /withdraw [post]
func (h Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var withdraw Withdraw

	accountID, _ := middleware.GetAccountID(r.Context())

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&withdraw)
	if err != nil {
		w.WriteHeader(r.Response.StatusCode)
		return
	}
	//mudar para account
	err = h.account.Withdraw(accountID, withdraw.Amount)
	if err != nil {
		w.WriteHeader(r.Response.StatusCode)
		return
	}

}
