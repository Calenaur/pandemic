package handler

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"github.com/Calenaur/pandemic/handler/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (h *Handler) helloTester(c echo.Context) error {

	username := c.FormValue("username")
	return c.JSON(http.StatusOK, username)

}

func (h *Handler) userbyid(e echo.Context) error {
	rowid := e.Param("id")
	id, err := strconv.ParseInt(rowid, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}
	user, err := h.us.GetByID(id)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}

	return e.JSON(http.StatusOK, user)
}

func (h *Handler) loginHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	// TODO No need to check input requirements for login?
	// err := inputRequirements(username, password)
	// if err != nil {
	// 	return response.MessageHandler(err, "", e)
	// }
	user, err := h.us.UserLogin(username, password)

	if err != nil {
		return response.MessageHandler(err, "", e)
	}

	if user == nil {
		return response.MessageHandler(err, "", e)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tok, err := token.SignedString([]byte(h.cfg.Token.Key))

	if err != nil {
		return response.MessageHandler(err, "UnhandledtokenError", e)
	}

	return e.JSON(http.StatusOK, map[string]string{
		"token": tok,
	})

}

func (h *Handler) signupHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	err := inputRequirements(username, password)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}

	err = h.us.UserSignup(username, password)

	if err != nil {
		return response.MessageHandler(err, "", e)
	}
	return response.MessageHandler(err, "User created", e)
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

func inputRequirements(username string, password string) error {
	var validUsername = regexp.MustCompile(`^([A-Za-z0-9]){2,16}$`)

	if !(len(password) >= 8 && len(password) <= 64) {
		return errors.New("Password length must be between 8 and 64 characters")
	}
	if !(len(username) >= 2 && len(username) <= 16) {
		return errors.New("Username length must be between 2 and 16 characters")
	}

next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"numeric":    {unicode.Number, unicode.Digit},
	} {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		// fmt.Printf("password must have at least one %s character", name)
		// fmt.Println()
		return errors.New("password must have at least one " + name + " character")
	}

	if !validUsername.MatchString(username) {
		return errors.New("Username can not have special characters")
	}
	return nil
}
