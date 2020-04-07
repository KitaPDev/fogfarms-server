package jwt

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/auth/sign_in", SignIn).
		Methods("POST").
		Schemes("http")
	router.HandleFunc("/auth/sign_out", SignOut).
		Methods("GET").
		Schemes("http")

	return router
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	AuthenticateSignIn(w, r)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	InvalidateToken(w)
}
