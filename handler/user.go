package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (h *Handler) helloTester(c echo.Context) error {

	username := c.FormValue("username")
	return c.JSON(http.StatusOK, username)

}

func (h *Handler) userbyid(e echo.Context) error {
	rowid := e.Param("id")
	id, inputErr := strconv.ParseInt(rowid, 10, 64)
	if inputErr != nil {
		e.JSON(CODE_ERROR_INVALID_ARGUMENTS, inputErr)
	}
	user, requestErr := h.us.GetByID(id)
	if requestErr != nil {
		e.JSON(CODE_ERROR_INTERNAL_SERVER_ERROR, requestErr)
	}

	return e.JSON(CODE_OK, user)

}

func (h *Handler) loginHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")
	user, err := h.us.UserLogin(username, password)

	if err != nil {
		e.JSON(CODE_ERROR_INTERNAL_SERVER_ERROR, err)
	}

	if user == nil {
		return e.JSON(http.StatusUnauthorized, echo.ErrUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tok, err := token.SignedString([]byte("جامعة هانزه العلوم تطبيقية"))

	if err != nil {
		return e.JSON(CODE_ERROR_INTERNAL_SERVER_ERROR, "something went wrong!")
	}

	return e.JSON(CODE_OK, map[string]string{
		"token": tok,
	})

}

func (h *Handler) signupHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	err := h.us.UserSignup(username, password)

	if err != nil {
		return e.JSON(http.StatusForbidden, "Duplicate entry")
	}
	return e.JSON(http.StatusCreated, "User created")
}

func accessible(c echo.Context) error {
	return c.JSON(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.JSON(http.StatusOK, map[string]string{
		"name": name,
	})
}
