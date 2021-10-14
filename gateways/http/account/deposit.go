package account

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/account"
	middleware "github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type DepositInput struct {
	Amount float64 `json: "amount"`
}

func ToDeposit(serv *mux.Router, usecase account.UseCase, log *logrus.Entry) *Handler {
	h := &Handler{
		account: usecase,
		log:     log,
	}

	serv.HandleFunc("/deposit", h.Deposit).Methods("Post")

	return h
}

// Deposit godoc
// @Summary Make a deposit
// @Description Make a deposit from an authenticated user
// @Param Body body DepositInput true "Body"
// @Accept  json
// @Produce  json
// @Success 201 "Created"
// @Failure 500 "Unable to read/write json"
// @Failure 401 "Invalid or missing token"
// @Param Authorization header string true "Bearer"
// @Router /deposit [post]
func (h Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	var deposit DepositInput
	usr, ok := middleware.GetCPF(r.Context())
	if !ok || usr == "" {
		h.log.Errorln("failed to authenticate user")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deposit)
	if err != nil {
		h.log.WithError(err).Errorln("unable to read json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	depositOutput, err := h.account.Deposit(usr, deposit.Amount)

	if err != nil {
		h.log.WithError(err).Errorln("failed to create deposit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(depositOutput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
