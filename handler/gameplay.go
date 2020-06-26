package handler

import (
	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
	"net/http"
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

//func (h *Handler) getAvailableMedicationsHandler(c echo.Context) error {
//	id, _, _ := getUserFromToken(c)
//
//	medications, err := h.us.GetDiseasesList(id)
//	if err != nil {
//		return response.MessageHandler(err, "", c)
//	}
//
//	return c.JSON(http.StatusOK, medications)
//}
