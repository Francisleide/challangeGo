package account

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h Handler) List_all_accounts(w http.ResponseWriter, r *http.Request) {
	accounts := h.account.List_all_accounts()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		log.Fatal(err)
	}

}
