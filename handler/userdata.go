package handler

import (
	"net/http"
	"strconv"

	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
)

// FormValue diseaseid
func (h *Handler) setUserDiseaseHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	disease := c.FormValue("diseaseid")
	diseaseid, err := strconv.Atoi(disease)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.ud.SetUserDisease(userid, diseaseid)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "Disease added to user", c)
}

func (h *Handler) getUserDiseaseHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	diseases, err := h.ud.GetUserDisease(userid)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, diseases)
}

// FIXME ID to user_disease
func (h *Handler) updateUserDiseaseHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)
	diseaseid := 0
	id := 0

	err := h.ud.UpdateUserDisease(userid, diseaseid, id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "User disease has been updated", c)
}

// FormValue tier
func (h *Handler) setUserTierHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	tier := c.FormValue("tier")
	tierid, err := strconv.Atoi(tier)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.ud.SetUserTier(userid, tierid)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "Tier added to user", c)
}

func (h *Handler) getUserTierHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	tiers, err := h.ud.GetUserTier(userid)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, tiers)
}

// FIXME ID to user_tier
func (h *Handler) updateUserTierHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)
	tier := c.FormValue("tier")
	tierInt, err := strconv.Atoi(tier)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.ud.UpdateUserTier(userid, tierInt)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "User tier has been updated", c)
}

// FormValue event
func (h *Handler) setUserEventHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	event := c.FormValue("event")
	eventid, err := strconv.Atoi(event)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.ud.SetUserEvent(userid, eventid)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "Event added to user", c)
}

func (h *Handler) getUserEventHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	events, err := h.ud.GetUserEvent(userid)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, events)
}

// FIXME add ID to user_event
func (h *Handler) updateUserEventHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)
	event := 0

	err := h.ud.UpdateUserEvent(userid, event)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "User event has been updated", c)
}

// FormValue researcher, researchername
func (h *Handler) setUserResearcherHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	researcher := c.FormValue("researcher")
	name := c.FormValue("researchername")

	researcherid, err := strconv.Atoi(researcher)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	err = h.ud.SetUserResearcher(userid, researcherid, name)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "Researcher added to user", c)
}

func (h *Handler) getUserResearcherHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	researchers, err := h.ud.GetUserResearcher(userid)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, researchers)
}

// FormValue oldname, newname
func (h *Handler) updateUserResearcherHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	oldName := c.FormValue("oldname")
	newName := c.FormValue("newname")

	err := h.ud.UpdateUserResearcher(userid, oldName, newName)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "Researcher name updated", c)
}

// FormValue researchername, traitname
func (h *Handler) setUserResearcherTraitHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	researcherName := c.FormValue("researchername")
	traitName := c.FormValue("traitname")

	err := h.ud.SetUserResearcherTrait(userid, researcherName, traitName)

	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return response.MessageHandler(err, "Researcher trait added to user researcher", c)
}

func (h *Handler) getUserResearcherTraitHandler(c echo.Context) error {
	userid, _, _ := getUserFromToken(c)

	researcherTraits, err := h.ud.GetUserResearcherTrait(userid)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	return c.JSON(http.StatusOK, researcherTraits)
}
