package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) helloTester(c echo.Context) error {

	return c.JSON(http.StatusOK, "HelloBud")

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

// func (h *Handler) loginHandler(e echo.Context) error {
// 	username := e.Param("username")
// 	password := e.Param("password")

// 	user, requestErr := h.us.GetByID()
// 	if requestErr != nil {
// 		e.JSON(CODE_ERROR_INTERNAL_SERVER_ERROR, requestErr)
// 	}

// 	return e.JSON(CODE_OK, user)

// }

// func (h *Handler) createUser(c echo.Context) error {

// }
