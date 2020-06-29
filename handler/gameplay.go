package handler

import (
	"net/http"

	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
)

func (h *Handler) medicationResearchHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	medication := c.FormValue("medication")
	err := h.us.ResearchMedication(id, medication)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, "Medication researched")
}
