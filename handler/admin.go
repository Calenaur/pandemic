package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
)

func (h *Handler) userbyid(e echo.Context) error {
	rowid := e.Param("id")
	_, _, accesslevel := getUserFromToken(e)
	accesslevelInt, err := strconv.ParseInt(accesslevel, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}
	// id, err := strconv.ParseInt(rowid, 10, 64)
	// if err != nil {
	// 	return response.MessageHandler(err, "", e)
	// }
	if accesslevelInt > 99 {
		user, err := h.us.GetByID(rowid)
		if err != nil {
			return response.MessageHandler(err, "", e)
		}
		return e.JSON(http.StatusOK, user)
	}
	err = errors.New("Restricted access")
	return response.MessageHandler(err, "", e)
}

func (h *Handler) listAll(e echo.Context) error {
	size := "10"
	_, _, accesslevel := getUserFromToken(e)
	accesslevelInt, err := strconv.ParseInt(accesslevel, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}
	page := e.QueryParam("page")
	offset, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}

	if e.QueryParam("size") != "" {
		size = e.QueryParam("size")
	}

	sizeInt, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", e)
	}

	if accesslevelInt > 99 {
		users, err := h.us.ListAll((offset-1)*sizeInt, sizeInt)
		if err != nil {
			return response.MessageHandler(err, "", e)
		}
		return e.JSON(http.StatusOK, users)
	}
	err = errors.New("Restricted access")
	return response.MessageHandler(err, "", e)
}

// Restricted admin access !TODO
func restricted(c echo.Context) error {
	userid, username, accesslevel := getUserFromToken(c)
	accesslevelInt, err := strconv.ParseInt(accesslevel, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	if accesslevelInt > 99 {
		return c.JSON(http.StatusOK, map[string]string{
			"name":        username,
			"id":          userid,
			"accesslevel": accesslevel,
		})
	}
	err = errors.New("Restricted access")

	return response.MessageHandler(err, "", c)

}

func (h *Handler) deleteUserByidHandler(c echo.Context) error {
	_, _, accesslevel := getUserFromToken(c)
	accesslevelInt, err := strconv.ParseInt(accesslevel, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	userId := c.FormValue("userid")
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	if accesslevelInt > 99 {
		err := h.us.DeleteUser(userId)
		if err != nil {
			return response.MessageHandler(err, "", c)
		}
	}

	return c.JSON(http.StatusOK, "User Deleted successfully")
}

func (h *Handler) makeUserAdminHandler(c echo.Context) error {
	_, _, accesslevel := getUserFromToken(c)
	accesslevelInt, err := strconv.ParseInt(accesslevel, 10, 64)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}
	userId := c.FormValue("userid")
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	if accesslevelInt > 99 {
		err := h.us.MakeUserAdmin(userId)
		if err != nil {
			return response.MessageHandler(err, "", c)
		}
	}

	return c.JSON(http.StatusOK, "User is now admin")
}
