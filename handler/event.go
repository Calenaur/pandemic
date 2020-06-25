package handler

import (
	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *Handler) getEventsHandler(c echo.Context) error {
	events, err := h.es.GetEvents()
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, events)

}

func (h *Handler) getEventByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	event, err := h.es.GetByID(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, event)
}
