package handler

import (
	"net/http"
	"strconv"

	"github.com/Calenaur/pandemic/handler/response"
	"github.com/Calenaur/pandemic/model"
	"github.com/labstack/echo"
)

func (h *Handler) getMedicationsHandler(c echo.Context) error {
	medications, err := h.ms.GetMedications()
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, medications)
}

func (h *Handler) getMedicationByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	medicationTrait, err := h.ms.GetTraitByID(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, medicationTrait)
}

func (h *Handler) medicationResearchHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	medication := c.FormValue("medication")
	err := h.us.ResearchMedication(id, medication)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, "Medication researched")
}

func (h *Handler) addUserMedicationAndTraits(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	req := c.Request();
	if err := req.ParseForm(); err != nil {
		return response.MessageHandler(err, "Couldn't parse request", c)
	}

	var medication string
	var traits []string

	if val, ok := req.PostForm["medication"]; ok {
		for _, s := range val {
			medication = s
		}
	}

	if val, ok := req.PostForm["trait"]; ok {
		for _, s := range val {
			traits = append(traits, s)
		}
	}

	newid, err := h.ms.AddMedicationAndTraits(id, medication, traits)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, &model.IDResponse{ID: newid})
}
