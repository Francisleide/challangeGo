package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/account"
	"github.com/francisleide/ChallangeGo/gateways/http/middlware"
	"github.com/gorilla/mux"
)

type DepositeHandler struct {
	account account.UseCase
}

type DepositeInput struct {
	Amount float64 `json: "amount"`
}

func ToDeposite(serv *mux.Router, usecase account.UseCase) *DepositeHandler {
	h := &DepositeHandler{
		account: usecase,
	}
	serv.HandleFunc("/deposite", h.Deposite).Methods("Post")

	return h
}

// ShowAccount godoc
// @Summary Make a deposite
// @Description Make a deposite from an authentic account
// @Param Body body Deposite true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "x-request-id"
// @Router /deposite [post]
func (h DepositeHandler) Deposite(w http.ResponseWriter, r *http.Request) {
	var deposite DepositeInput
	usr, ok := middlware.GetAccountID(r.Context())
	if !ok {
		log.Fatal("Usuário não autenticado: ")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&deposite)
	if err != nil {
		log.Fatal("Não consegui ler o body: ", err)
	}
	//mudar para account
	h.account.Deposite(usr, deposite.Amount)

	if err != nil {
		log.Fatal(err)
	}

}
