package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

func (h *Handler) DebugHandler(c echo.Context) error {
	user, err := h.us.GetByID(1)
	if (err != nil) {
		data := echo.Map {
			"title": "Debug",
			"message": err.Error(),
		}
		return c.Render(http.StatusOK, "index", data)
	}

	return c.JSON(http.StatusOK, user)
}