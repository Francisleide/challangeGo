package autenticationoperations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/gateways/middlware"
)

type Deposite struct {
	Ammount float64 `json: "ammount"`
}

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
		log.Fatal("Não consegui ler o body")
	}

	h.autentic.Deposite(usr, deposite.Ammount)

	if err != nil {
		log.Fatal(err)
	}
	
}
