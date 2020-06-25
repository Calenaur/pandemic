package handler

import (
	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *Handler) getMedicationsHandler(c echo.Context) error {
	medications, err := h.ms.GetMedications()
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, medications)
}

func (h *Handler) getMedicationByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	medication, err := h.ms.GetByID(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, medication)
}

func (h *Handler) getMedicationTraitsHandler(c echo.Context) error {
	medicationTraits, err := h.ms.GetTraits()
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, medicationTraits)

}

func (h *Handler) getMedicationTraitByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	medicationTrait, err := h.ms.GetTraitByID(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, medicationTrait)
}