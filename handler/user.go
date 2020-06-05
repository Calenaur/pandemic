package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"unicode"

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
		log.Fatal(err)
		return e.JSON(http.StatusBadRequest, "Id not an int.")
	}
	user, err := h.us.GetByID(id)
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusInternalServerError, "Can't find id.")
	}

	return e.JSON(http.StatusOK, user)
}

func (h *Handler) loginHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	if !inputRequirements(username, password) {
		return e.JSON(http.StatusForbidden, "Username or password does not meet requirements.")
	}
	user, err := h.us.UserLogin(username, password)

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return e.JSON(http.StatusUnauthorized, "Token malformed.")
	}

	return e.JSON(http.StatusOK, map[string]string{
		"token": tok,
	})

}

func (h *Handler) signupHandler(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	if !inputRequirements(username, password) {
		return e.JSON(http.StatusForbidden, "Username or password does not meet requirements.")
	}

	err := h.us.UserSignup(username, password)

	if err != nil {
		log.Fatal(err)
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

func inputRequirements(username string, password string) bool {
	var validUsername = regexp.MustCompile(`^([A-Za-z0-9]){2,16}$`)

	if !(len(password) >= 8 && len(password) <= 64) {
		return false
	}

next:
	for _, classes := range map[string][]*unicode.RangeTable{
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
		return false
	}

	return validUsername.MatchString(username)
}
