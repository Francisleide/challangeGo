package autenticationoperations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/gateways/http/middlware"
)

type Withdraw struct {
	Amount float64 `json: "amount"`
}

// ShowAccount godoc
// @Summary Make a Withdraw
// @Description Make a Withdraw from an authentic account
// @Param Body body Withdraw true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "x-request-id"
// @Router /withdraw [post]
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
	ok := h.autentic.WithDraw(accountId, withdraw.Amount)
	if !ok {
		log.Panic("Saldo insuficiente!")
	}

}
