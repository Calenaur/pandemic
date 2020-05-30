package handler

import (
	"net/http"
	"strconv"

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

	return e.JSON(CODE_OK, user)

}

func (h *Handler) createUser(c echo.Context) error {
	return nil
}
