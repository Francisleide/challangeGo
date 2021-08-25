package gateways

import (
	"fmt"
	"log"
	"net/http"

	// gin-swagger middleware
	// swagger embed files
	"github.com/francisleide/ChallengeGo/domain/account/usecase"
	acc "github.com/francisleide/ChallengeGo/domain/account/usecase"
	a "github.com/francisleide/ChallengeGo/domain/auth/usecase"
	tr "github.com/francisleide/ChallengeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/francisleide/ChallengeGo/gateways/http/auth"
	middleware "github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/francisleide/ChallengeGo/gateways/http/transfer"
	"github.com/gorilla/mux"
	http_swagger "github.com/swaggo/http-swagger"
)

type Api struct {
	account  acc.AccountUc
	transfer tr.TransferUc
	auth     a.AuthUc
}

func NewApi(accountUC usecase.AccountUc, transferUC tr.TransferUc, authorization a.AuthUc) *Api {
	return &Api{
		account:  accountUC,
		transfer: transferUC,
		auth:     authorization,
	}
}

func (api Api) Run(host string, port string) {
	r := mux.NewRouter()

	authenticatedRoute := r.PathPrefix("").Subrouter()
	unauthenticatedRoute := r.PathPrefix("").Subrouter()
	account.Accounts(unauthenticatedRoute, api.account)
	transfer.Transfer(authenticatedRoute, api.transfer)
	account.ToDeposit(authenticatedRoute, api.account)
	account.ToWithdraw(authenticatedRoute, api.account)
	auth.Auth(r, api.auth)
	fmt.Println("Executing Run() with:  ", host, port)
	r.PathPrefix("/docs/swagger").Handler(http_swagger.WrapHandler).Methods(http.MethodGet)

	authenticatedRoute.Use(middleware.Authorize)
	endpoint := fmt.Sprintf("%s:%s", host, port)
	serv := &http.Server{
		Handler: r,
		Addr:    endpoint,
	}

	log.Fatal(serv.ListenAndServe())

}
