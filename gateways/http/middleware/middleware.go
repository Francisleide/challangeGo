package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/francisleide/ChallengeGo/domain/auth/usecase"
	"github.com/sirupsen/logrus"
)

type AuthContextKey string

var ContextID = AuthContextKey("cpf")

func ValidateToken(next http.Handler) http.Handler {
	var x *jwt.Token

	log := logrus.New()
	logEntry := logrus.NewEntry(log)
	var err error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		strArr := strings.Split(token, " ")
		if len(strArr) == 2 {
			x, err = VerifyToken(strArr[1])
			if err != nil {
				logEntry.Errorf("invalid token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		accountID := usecase.Authentication(x)

		ctx := context.WithValue(r.Context(), ContextID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func VerifyToken(tokenstr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetCPF(ctx context.Context) (string, bool) {
	tokenStr, ok := ctx.Value(ContextID).(string)

	return tokenStr, ok
}
