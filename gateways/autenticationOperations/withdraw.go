package autenticationoperations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/gateways/middlware"
)

type Withdraw struct {
	Ammount float64 `json: "ammount"`
}

func (h Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var withdraw Withdraw

	accountId, _ := middlware.GetAccountID(r.Context())
	fmt.Println(accountId)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&withdraw)
	if err != nil {
		log.Fatal(err)
	}
	ok:=h.autentic.WithDraw(accountId, withdraw.Ammount)
	if !ok{
		log.Panic("Saldo insuficiente!")
	}
	fmt.Println(withdraw.Ammount)
	fmt.Println("Conta : ", accountId)

}
