package autenticationoperations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/gateways/http/middlware"
)

type Deposite struct {
	Amount float64 `json: "ammount"`
}

// ShowAccount godoc
// @Summary Make a deposite
// @Description Make a deposite from an authentic account
// @Param Body body Deposite true "Body"
// @Accept  json
// @Produce  json
// @Header 201 {string} Token "x-request-id"
// @Router /deposite [post]
func (h Handler) Deposite(w http.ResponseWriter, r *http.Request) {
	var deposite Deposite
	usr, ok := middlware.GetAccountID(r.Context())
	if !ok {
		log.Fatal("Usuário não autenticado: ")
	}
	fmt.Println("Usuario: ", usr)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&deposite)
	if err != nil {
		log.Fatal("Não consegui ler o body: ", err)
	}

	h.autentic.Deposite(usr, deposite.Amount)

	if err != nil {
		log.Fatal(err)
	}

}
