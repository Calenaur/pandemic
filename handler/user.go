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
	claims["access"] = user.AccessLevel
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

	err := usernameRequirements(username)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}
	err = passwordRequirements(password)
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

// Restricted admin access !TODO
func restricted(c echo.Context) error {
	userid, username, accesslevel := getUserFromToken(c)

	if accesslevel < "1" {
		return c.JSON(http.StatusOK, map[string]string{
			"name":        username,
			"id":          userid,
			"accesslevel": accesslevel,
		})
	}

	return c.JSON(http.StatusBadRequest, "yOu aReN't AlLod Her")

}

// Allow the user to change his/her name
func (h *Handler) changeNameHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	newname := c.FormValue("newname")

	err := usernameRequirements(newname)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.us.ChangeUserName(id, newname)
	if err != nil {
		return response.MessageHandler(err, "Name could not be changed", c)
	}
	return response.MessageHandler(err, "Name updated successfully", c)
}

// Allow the user to Change their password
func (h *Handler) changePasswordHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	newPassword := c.FormValue("newpassword")

	err := passwordRequirements(newPassword)
	if err != nil {
		return response.MessageHandler(err, "", c)

	}

	err = h.us.ChangeUserPassword(id, newPassword)
	if err != nil {
		return response.MessageHandler(err, "Password could not be changed", c)
		// return c.JSON(http.StatusBadRequest, "Password couldn't be changed please try again")
	}
	return response.MessageHandler(err, "Password updated successfully", c)
	// return c.JSON(http.StatusOK, "Passowrd changed successfully")
}

// Allow the user to delete their own account
func (h *Handler) deleteAccountHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	err := h.us.DeleteAccount(id)

	if err != nil {
		return response.MessageHandler(err, "UnknownError", c)
		// return c.JSON(http.StatusBadRequest, "Account couldn't be deleted")
	}
	return response.MessageHandler(err, "Account deleted successfully", c)
	// return c.JSON(http.StatusOK, "Account deleted successfully")
}

// this function returns the user id and username from the token,
// user id and name here are strings to make it easier to use them
// in SQL querries
func getUserFromToken(c echo.Context) (string, string, string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["name"].(string)
	rowid := claims["sub"].(float64)

	accesslevel := claims["access"].(float64)

	stringID := fmt.Sprintf("%g", rowid)
	stringLevel := fmt.Sprintf("%g", accesslevel)

	fmt.Println(stringID)
	fmt.Println(stringLevel)

	return stringID, username, stringLevel
}

func usernameRequirements(username string) error {
	var validUsername = regexp.MustCompile(`^([A-Za-z0-9]){2,16}$`)

	if !(len(username) >= 2 && len(username) <= 16) {
		return errors.New("Username length must be between 2 and 16 characters")
	}

	if !validUsername.MatchString(username) {
		return errors.New("Username can not have special characters")
	}
	return nil
}

func passwordRequirements(password string) error {
	if !(len(password) >= 8 && len(password) <= 64) {
		return errors.New("Password length must be between 8 and 64 characters")
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
	return nil
}
