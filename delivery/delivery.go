package delivery

import (
	"casbin/models"
	"casbin/usecase"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

type (
	newDeliveryHandler struct {
		Cs usecase.Usecase
	}
	M struct {
		Message string `json:"message"`
	}
)

func (h *newDeliveryHandler) GetPerson(c echo.Context) error {
	ctx := c.Request().Context()
	var data, err = h.Cs.GetPersons(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *newDeliveryHandler) Login(c echo.Context) error {
	username, password, ok := c.Request().BasicAuth()
	if !ok {
		return c.JSON(http.StatusBadRequest, M{"Invalid Login"})
	}
	ctx := c.Request().Context()
	user, err := h.Cs.GetUser(ctx, username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{err.Error()})
	}

	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    models.APP_NAME,
			ExpiresAt: time.Now().Add(models.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(
		models.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(models.JWT_SIGNATURE_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{err.Error()})
	}
	c.Response().Header().Set("x-token", signedToken)
	c.Response().Header().Set("content-type", "application/json")
	return c.JSON(http.StatusOK, M{"Sukses Login"})
}

func NewHandler(e *echo.Echo, u usecase.Usecase) {
	handler := &newDeliveryHandler{Cs: u}
	e.GET("/person", handler.GetPerson)
	e.GET("/login", handler.Login)
	e.POST("/person", handler.GetPerson)
}
