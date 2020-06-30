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

	balance := strconv.Itoa(user.Balance)
	tier := strconv.Itoa(user.Tier)

	return e.JSON(http.StatusOK, map[string]string{
		"token":   tok,
		"balance": balance,
		"tier":    tier,
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

// Get user details like Balance and Manufacture
func (h *Handler) getUserDetailsHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	username, accesslevel, tier, balance, err := h.us.GetUserDetails(id)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id":          id,
		"username":    username,
		"accesslevel": accesslevel,
		"tier":        tier,
		"balance":     balance,
	})
}

func (h *Handler) getFriendsHandler(c echo.Context) error {
	id, name, _ := getUserFromToken(c)

	friends, err, _ := h.us.ShowFriends(id, name)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, friends)
}

func (h *Handler) pendingFriendsHandler(c echo.Context) error {
	id, name, _ := getUserFromToken(c)

	friends, err, _ := h.us.ShowPendingFriends(id, name)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, friends)
}

// Update the user Balance
func (h *Handler) updateBalanceHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	newBalance := c.FormValue("newbalance")

	err := h.us.UpdateBalance(id, newBalance)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, "Balance Changed successfully")
}

func (h *Handler) updateDeviceHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	device := c.FormValue("device")

	err := h.us.UpdateDevice(id, device)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, "device Changed successfully")
}

func (h *Handler) getDeviceHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	device, err := h.us.GetDevice(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, device)
}

// Allow the user to change his/her name
func (h *Handler) changeNameHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	// fmt.Println(id)

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

func (h *Handler) sendFriendRequestHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	friend := c.FormValue("friend")

	err := h.us.SendFriendRequest(id, friend)
	if err != nil {
		return response.MessageHandler(err, "No user by this name was found", c)
		// return c.JSON(http.StatusBadRequest, "Account couldn't be deleted")
	}

	//return c.JSON(http.StatusOK, "Friend request send successfully")
	return response.MessageHandler(err, "Friend request sent", c)
}

func (h *Handler) responseFriendRequestHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	friend := c.FormValue("friend")
	rowResponse := c.FormValue("response")

	respons, erro := strconv.ParseInt(rowResponse, 10, 64)
	if erro != nil {
		return response.MessageHandler(erro, "Invalid response", c)
	}

	// Typo intended
	err := h.us.RespondFriendRequest(id, friend, respons)
	if err != nil {
		return response.MessageHandler(err, "Something went wrong", c)
		// return c.JSON(http.StatusBadRequest, "Account couldn't be deleted")
	}

	return c.JSON(http.StatusOK, "Responded successfully")
}

func (h *Handler) deleteFriendHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	friend := c.FormValue("friend")

	// Typo intended
	err := h.us.DeleteFriend(id, friend)
	if err != nil {
		return response.MessageHandler(err, "Something went wrong", c)
		// return c.JSON(http.StatusBadRequest, "Account couldn't be deleted")
	}

	return response.MessageHandler(err, "Friend deleted", c)
}

func (h *Handler) giftFriendHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	friend := c.FormValue("friend")
	balance := c.FormValue("balance")

	// Typo intended
	err := h.us.SendFriendBalance(id, friend, balance)
	if err != nil {
		return response.MessageHandler(err, "Something went wrong", c)
		// return c.JSON(http.StatusBadRequest, "Account couldn't be deleted")
	}

	//return c.JSON(http.StatusOK, "Balance send su")
	return response.MessageHandler(err, "Balance sent", c)
}

// this function returns the user id and username from the token,
// user id and name here are strings to make it easier to use them
// in SQL querries
func getUserFromToken(c echo.Context) (string, string, string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["name"].(string)
	rowid := claims["sub"].(string)

	accesslevel := claims["access"].(float64)

	// stringID := fmt.Sprintf("%g", rowid)
	stringLevel := fmt.Sprintf("%g", accesslevel)

	// fmt.Println(stringID)
	// fmt.Println(stringLevel)

	return rowid, username, stringLevel
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

func (h *Handler) getUserMedicationsHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	userMedications, err := h.us.GetUserMedications(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, userMedications)
}

func (h *Handler) getUserMedicationByIDHandler(c echo.Context) error {
	userID, _, _ := getUserFromToken(c)
	userMedicationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	userMedications, err := h.us.GetUserMedicationByID(userID, userMedicationID)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, userMedications)
}
