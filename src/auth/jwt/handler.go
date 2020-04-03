package jwt

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("POST", "/auth/sign_in", AuthenticateSignIn)
	router.HandlerFunc("GET", "/auth/refresh", Refresh)
	router.HandlerFunc("GET", "/auth/sign_out", SignOut)

	return router
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	RefreshToken(w, r)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	InvalidateToken(w)
}
