package middlewares

import (
	"casbin/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
)

type M map[string]interface{}
type Enforcer struct {
	Enforcer *casbin.Enforcer
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		tokenToString := r.Header.Get("x-token")

		token, err := jwt.Parse(tokenToString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != models.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}
			return models.JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(M{
				"message": err.Error(),
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(M{
				"message": err.Error(),
			})
			return
		}
		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (e *Enforcer) Enforce(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get url path
		path := r.URL.Path

		// if url path = /login, skip
		if path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		// get token from header
		token := r.Header.Get("x-token")

		// get method. POST, GET, PUT, DELETE
		method := r.Method

		// decode and get role from jwt
		user, err := decodeToken(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(M{
				"message": err.Error(),
			})
			return
		}

		result, err := e.Enforcer.Enforce(user, path, method)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(M{
				"message": err.Error(),
			})
			return
		}

		if result {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(M{"message": "anda tidak memiliki akses"})
	})
}

func decodeToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(models.JWT_SIGNATURE_KEY), nil
		},
	)
	if err != nil || !token.Valid {
		return "", err
	}
	return claims["role"].(string), nil
}
