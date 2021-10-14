package auth

import (
	"encoding/json"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/auth"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Login struct {
	CPF    string
	Secret string
}

type Handler struct {
	auth auth.UseCase
	log  *logrus.Entry
}

func Auth(serv *mux.Router, usecase auth.UseCase, log *logrus.Entry) *Handler {
	h := &Handler{
		auth: usecase,
		log:  log,
	}

	serv.HandleFunc("/login", h.Authentication).Methods("Post")

	return h
}

// Authentication godoc
// @Summary Login
// @Description Takes the CPF and password of a user, if they are correct, a token is generated
// @Param Body body Login true "Body"
// @Accept  json
// @Produce  json
// @Success 200 "Ok"
// @Failure 400 "Unable to read/write json"
// @Failure 500 "Failed to create token"
// @Router /login [post]
func (h Handler) Authentication(w http.ResponseWriter, r *http.Request) {
	var login Login
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&login)
	if err != nil {
		h.log.WithError(err).Errorln("unable to read json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := h.auth.CreateToken(login.CPF, login.Secret)
	if err != nil {
		h.log.WithError(err).Errorln("failed to create token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		h.log.WithError(err).Errorln("unable to write json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
