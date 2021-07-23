package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/auth"
	"github.com/gorilla/mux"
)

type Login struct {
	Cpf    string `json: "cpf"`
	Secret string `json: "secret"`
}

type Handler struct {
	auth auth.UseCase
}

func Auth(serv *mux.Router, usecase auth.UseCase) *Handler {
	h := &Handler{
		auth: usecase,
	}

	serv.HandleFunc("/auth", h.Authorization).Methods("Post")

	return h
}

// ShowAccount godoc
// @Summary Get a Auth
// @Description It takes a token to authenticate yorself to the application
// @Param Body body Login true "Body"
// @Accept  json
// @Produce  json
// @Router /auth [post]
func (h Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	var login Login
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&login)
	fmt.Println("cpf: ", login.Cpf)
	fmt.Println("senha: ", login.Secret)
	if err != nil {
		log.Fatal(err)
	}
	token, err := h.auth.CreateToken(login.Cpf, login.Secret)
	err = json.NewEncoder(w).Encode(token)
}
