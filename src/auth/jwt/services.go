package jwt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/gddo/httputil/header"
	"github.com/labstack/gommon/log"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	//jwtKey := os.Getenv("SECRET_KEY_JWT")
	jwtKey := "s"
	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Print("True")
	return
}

func AuthenticateSignIn(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username string
		Password string
	}
	var testdata Input
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&testdata)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Printf("%+v", testdata)
	username := testdata.Username
	password := testdata.Password

	// exists, _ := user.Exists(username)
	// if !exists {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	log.Fatal(io.WriteString(w, `{"error":"user_not_found"}"`))
	// 	return
	// }

	valid := user.ValidateUserA(username, password)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	GenerateToken(username, w)
}

func GenerateToken(username string, w http.ResponseWriter) {
	//jwtKey := os.Getenv("SECRET_KEY_JWT")
	jwtKey := "s"
	fmt.Printf("%+v", jwtKey)
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(io.WriteString(w, `{"error":"token_generation_failed"`))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "jwtToken",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	GenerateToken(claims.Username, w)
}

func InvalidateToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwtToken",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}

func generateAuthHeaderFromToken(token string) string {
	return fmt.Sprintf(bearerFormat, token)
}
