package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
)

func (h *Handler) getTierByIDHandler(c echo.Context) error {
	tier := c.Param("id")
	tierid, err := strconv.Atoi(tier)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	tierByID, err := h.ts.GetTierByID(tierid)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, tierByID)
}

func (h *Handler) getTierListHandler(c echo.Context) error {
	tiers, err := h.ts.GetTierList()
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, tiers)
}

// ADMIN
func (h *Handler) setTierHandler(c echo.Context) error {
	_, _, accesslevel := getUserFromToken(c)
	accesslevelInt, err := strconv.Atoi(accesslevel)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	tiername := c.FormValue("tiername")
	tiercolor := c.FormValue("tiercolor")

	if accesslevelInt > 99 {
		err := h.ts.SetTier(tiername, tiercolor)
		if err != nil {
			return response.MessageHandler(err, "", c)
		}
		return response.MessageHandler(err, "Tier added", c)
	}
	err = errors.New("Restricted access")

	return response.MessageHandler(err, "", c)
}

func (h *Handler) updateTierHandler(c echo.Context) error {
	_, _, accesslevel := getUserFromToken(c)
	accesslevelInt, err := strconv.Atoi(accesslevel)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	if accesslevelInt < 100 {
		err = errors.New("Restricted access")
		return response.MessageHandler(err, "", c)
	}
	id := c.Param("id")
	tiername := c.FormValue("tiername")
	tiercolor := c.FormValue("tiercolor")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	tier, err := h.ts.GetTierByID(idInt)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	if tiername == "" {
		tiername = tier.Name
	}
	if tiercolor == "" {
		tiercolor = tier.Color
	}

	err = h.ts.UpdateTier(idInt, tiername, tiercolor)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return response.MessageHandler(err, "Tier updated", c)
}

func (h *Handler) deleteTierHandler(c echo.Context) error {
	_, _, accesslevel := getUserFromToken(c)
	accesslevelInt, err := strconv.Atoi(accesslevel)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	if accesslevelInt < 100 {
		err = errors.New("Restricted access")
		return response.MessageHandler(err, "", c)
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.ts.DeleteTier(idInt)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return response.MessageHandler(err, "Tier deleted", c)
}
