package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	APP_NAME                  = "Golang RBAC with Casbin"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("Apasaja Boleh :p")
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
