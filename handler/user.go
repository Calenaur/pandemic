package handler

import (
	"fmt"
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
		return e.JSON(http.StatusBadRequest, "Wrong id provided.")
	}
	user, requestErr := h.us.GetByID(id)
	if requestErr != nil {
		return e.JSON(http.StatusInternalServerError, "Can't find id.")
	}

	return e.JSON(http.StatusOK, user)
}

func (h *Handler) loginHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	user, err := h.us.UserLogin(username, password)

	if err != nil {
		return e.JSON(http.StatusServiceUnavailable, "Database can not handle the request.")
	}

	if user == nil {
		return e.JSON(http.StatusUnauthorized, "Wrong username or password.")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tok, err := token.SignedString([]byte(h.cfg.Token.Key))

	if err != nil {
		return e.JSON(http.StatusUnauthorized, "Token malformed.")
	}

	return e.JSON(http.StatusOK, map[string]string{
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
	fmt.Println(claims)
	rowid := claims["sub"].(float64)

	sub := fmt.Sprintf("%g", rowid)

	fmt.Print(sub)
	return c.JSON(http.StatusOK, map[string]string{
		"name": name,
		"id":   sub,
	})
}

// Allow the user to change his/her name
func (h *Handler) changeNameHandler(c echo.Context) error {
	id, _ := getUserFromToken(c)

	newname := c.FormValue("newname")

	err := h.us.ChangeUserName(id, newname)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Name couldn't be changed please try again")
	}

	return c.JSON(http.StatusOK, "username changed successfully")
}

// Allow the user to Change their password
func (h *Handler) changePasswordHandler(c echo.Context) error {
	id, _ := getUserFromToken(c)

	newPassword := c.FormValue("newpassword")

	err := h.us.ChangeUserPassword(id, newPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Password couldn't be changed please try again")
	}

	return c.JSON(http.StatusOK, "Passowrd changed successfully")
}

// Allow the user to delete their own account
func (h *Handler) deleteAccountHandler(c echo.Context) error {
	id, _ := getUserFromToken(c)

	err := h.us.DeleteAccount(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Account couldn't be deleted")
	}

	return c.JSON(http.StatusOK, "Account deleted successfully")
}

// this function returns the user id and username from the token,
// user id and name here are strings to make it easier to use them
// in SQL querries
func getUserFromToken(c echo.Context) (string, string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["name"].(string)
	rowid := claims["sub"].(float64)

	stringID := fmt.Sprintf("%g", rowid)

	return stringID, username
}
