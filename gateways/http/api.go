package gateways

import (
	"fmt"
	"log"
	"net/http"

	// gin-swagger middleware
	// swagger embed files
	"github.com/francisleide/ChallangeGo/domain/account/usecase"
	acc "github.com/francisleide/ChallangeGo/domain/account/usecase"
	a "github.com/francisleide/ChallangeGo/domain/auth/usecase"
	tr "github.com/francisleide/ChallangeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallangeGo/gateways/http/account"
	"github.com/francisleide/ChallangeGo/gateways/http/auth"
	"github.com/francisleide/ChallangeGo/gateways/http/middlware"
	"github.com/francisleide/ChallangeGo/gateways/http/transfer"
	"github.com/gorilla/mux"
	http_swagger "github.com/swaggo/http-swagger"
)

///var c_account account.Handler

type Api struct {
	account  acc.AccountUc
	transfer tr.TransferUc
	auth     a.AuthUc
	//autentic aut.Autentic
}

func NewApi(acc usecase.AccountUc, transf tr.TransferUc, authorization a.AuthUc) *Api {
	return &Api{
		account:  acc,
		transfer: transf,
		auth:     authorization,
	}
}

func (api Api) Run(host string, port string) {
	r := mux.NewRouter()

	Auth := r.PathPrefix("").Subrouter()
	NoAuth := r.PathPrefix("").Subrouter()
	account.Accounts(NoAuth, api.account)
	transfer.Transfer(Auth, api.transfer)
	account.ToDeposite(Auth, api.account)
	account.ToWithdraw(Auth, api.account)
	auth.Auth(r, api.auth)
	fmt.Println("Executing Run() with:  ", host, port)
	//autenticationoperations.AutenticationOperations(Auth, api.autentic)
	r.PathPrefix("/docs/swagger").Handler(http_swagger.WrapHandler).Methods(http.MethodGet)

	Auth.Use(middlware.Authorize)
	endpoint := fmt.Sprintf("%s:%s", host, port)
	serv := &http.Server{
		Handler: r,
		Addr:    endpoint,
	}

	log.Fatal(serv.ListenAndServe())

}
