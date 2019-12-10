package middlewares

import (
	"casbin/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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

func (e *Enforcer) Enforce(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get url path
		path := c.Request().URL.Path

		// if url path = /login, skip
		if path == "/login" {
			return next(c)
		}

		// get token from header
		token := c.Request().Header.Get("x-token")

		// get method. POST, GET, PUT, DELETE
		method := c.Request().Method

		// decode and get role from jwt
		user, err := decodeToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, M{"message": err.Error()})
		}

		result, err := e.Enforcer.Enforce(user, path, method)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, M{"message": err.Error()})
		}

		if result {
			return next(c)
		}
		return c.JSON(http.StatusForbidden, M{"message": "Forbidden"})
	}
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
