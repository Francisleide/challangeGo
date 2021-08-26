package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallengeGo/domain/auth"
	"github.com/gorilla/mux"
)

type Login struct {
	CPF    string
	Secret string
}

type Handler struct {
	auth auth.UseCase
}

func Auth(serv *mux.Router, usecase auth.UseCase) *Handler {
	h := &Handler{
		auth: usecase,
	}

	serv.HandleFunc("/login", h.Authentication).Methods("Post")

	return h
}

// ShowAccount godoc
// @Summary Get a Auth
// @Description It takes a token to authenticate yourself to the application
// @Param Body body Login true "Body"
// @Accept  json
// @Produce  json
// @Router /login [post]
func (h Handler) Authentication(w http.ResponseWriter, r *http.Request) {
	var login Login
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&login)
	if err != nil {
		log.Fatal(err)
	}
	token, err := h.auth.CreateToken(login.CPF, login.Secret)
	err = json.NewEncoder(w).Encode(token)
}
