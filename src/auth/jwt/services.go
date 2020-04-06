package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/gddo/httputil/header"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) bool {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		if err == http.ErrNoCookie {
			msg := "Error: No Token Found"
			http.Error(w, msg, http.StatusUnauthorized)
			return false
		}

		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			msg := "Error: Invalid Signature"
			http.Error(w, msg, http.StatusUnauthorized)
			return false
		}

		msg := "Error: Failed to Parse Token"
		http.Error(w, msg, http.StatusUnauthorized)
		return false
	}

	if !token.Valid {
		msg := "Error: Invalid Token"
		http.Error(w, msg, http.StatusUnauthorized)
		return false
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 0*time.Second {

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			msg := "Error: Token Exceeded Timeout Limit"
			http.Error(w, msg, http.StatusUnauthorized)
			return false

		} else {
			GenerateToken(claims.Username, w)
		}

	}

	return true
}

func AuthenticateSignIn(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username string
		Password string
	}
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var credentials Input

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")

		if value != "application/json" {
			msg := "Error: Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		msg := "Failed to Decode"
		http.Error(w, msg, http.StatusBadRequest)
	}

	username := credentials.Username
	//password := credentials.Password

	exists, _ := user.ExistsByUsername(username)
	if !exists {
		msg := "Error: User Not Found"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	// exists, _ := user.ExistsByUsername(username)
	// if !exists {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	log.Fatal(io.WriteString(w, `{"error":"user_not_found"}"`))
	// 	return
	// }

	GenerateToken(username, w)
}

func GenerateToken(username string, w http.ResponseWriter) {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	expirationTime := time.Now().Add(10 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  expirationTime,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		msg := "Error: Failed to Generate Token"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwtToken",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

//func RefreshToken(w http.ResponseWriter, r *http.Request) {
//	jwtKey := os.Getenv("SECRET_KEY_JWT")
//
//	cookie, err := r.Cookie("jwtToken")
//	if err != nil {
//		if err == http.ErrNoCookie {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	tokenString := cookie.Value
//	claims := &Claims{}
//
//	token, err := jwt.ParseWithClaims(tokenString, claims,
//		func(token *jwt.Token) (interface{}, error) {
//			return jwtKey, nil
//		})
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	if !token.Valid {
//		w.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//
//	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	GenerateToken(claims.Username, w)
//}

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
