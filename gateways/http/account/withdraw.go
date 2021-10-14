package account

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	middleware "github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Withdraw struct {
	Amount float64 `json: "amount"`
}

func ToWithdraw(serv *mux.Router, usecase account.UseCase, log *logrus.Entry) *Handler {
	h := &Handler{
		account: usecase,
		log:     log,
	}
	serv.HandleFunc("/withdraw", h.Withdraw).Methods("Post")
	return h
}

// Withdraw godoc
// @Summary Make a Withdraw
// @Description Make a Withdraw from an authentic account
// @Param Body body Withdraw true "Body"
// @Accept  json
// @Produce  json
// 
// @Failure 401 "Invalid or missing token"
// @Param Authorization header string true "Bearer"
// @Router /withdraw [post]
func (h Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var withdraw Withdraw

	accountID, ok := middleware.GetCPF(r.Context())
	if !ok || accountID == "" {
		h.log.Errorln("failed to authenticate user")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&withdraw)
	if err != nil {
		h.log.WithError(err).Errorln("unable to read json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	withdrawOutput, err := h.account.Withdraw(accountID, withdraw.Amount)
	if err != nil {
		h.log.WithError(err).Errorln("failed to create withdraw")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(withdrawOutput)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		h.log.WithError(err).Errorln("unable to generate json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
