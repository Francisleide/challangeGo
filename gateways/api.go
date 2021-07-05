package gateways

import (
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/account/usecase"
	acc "github.com/francisleide/ChallangeGo/domain/account/usecase"
	aut "github.com/francisleide/ChallangeGo/domain/autenticOperations/usecase"
	a "github.com/francisleide/ChallangeGo/domain/auth/usecase"
	tr "github.com/francisleide/ChallangeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallangeGo/gateways/account"
	autenticationoperations "github.com/francisleide/ChallangeGo/gateways/autenticationOperations"
	"github.com/francisleide/ChallangeGo/gateways/auth"
	"github.com/francisleide/ChallangeGo/gateways/middlware"
	"github.com/francisleide/ChallangeGo/gateways/transfer"
	"github.com/gorilla/mux"
)

var c_account account.Handler

type Api struct {
	account  acc.AccountUc
	transfer tr.TransferUc
	auth     a.AuthUc
	autentic aut.Autentic
}

func NewApi(acc usecase.AccountUc, transf tr.TransferUc, autentic aut.Autentic, authorization a.AuthUc) *Api {
	return &Api{
		account:  acc,
		transfer: transf,
		autentic: autentic,
		auth:     authorization,
	}
}

func (a Api) Run(host string, port string) {
	r := mux.NewRouter()
	Auth := r.PathPrefix("").Subrouter()
	account.Accounts(r, a.account)
	transfer.Transfer(Auth, a.transfer)
	auth.Auth(r, a.auth)
	autenticationoperations.AutenticationOperations(Auth, a.autentic)

	Auth.Use(middlware.Authorize)
	endpoint := fmt.Sprintf("%s:%s", host, port)
	serv := &http.Server{
		Handler: r,
		Addr:    endpoint,
	}
	log.Fatal(serv.ListenAndServe())

}
