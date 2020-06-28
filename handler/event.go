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
		return response.MessageHandler(err, "This isn't working", c)
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

func (h *Handler) getMyEventsHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	events, err := h.es.GetMyEvents(id)
	if err != nil {
		return response.MessageHandler(err, "This isn't working", c)
	}

	return c.JSON(http.StatusOK, events)

}

func (h *Handler) subscribeToEventHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	event := c.FormValue("event")

	err := h.es.SubscribeToEvent(id, event)
	if err != nil {
		return response.MessageHandler(err, "Couldn't subscribe to event: "+event, c)
	}
	return response.MessageHandler(err, "subscribed to Event: "+event, c)
}

func (h *Handler) unSubscribeToEventHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)
	event := c.FormValue("event")

	err := h.es.UnSubscribeToEvent(id, event)
	if err != nil {
		return response.MessageHandler(err, "Couldn't subscribe to event: "+event, c)
	}
	return response.MessageHandler(err, "unsubscribed to Event: "+event, c)
}
