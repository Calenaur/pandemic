package handler

import (
	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/store"
	"github.com/labstack/echo"
)

const CODE_OK = 200
const CODE_ERROR_INVALID_ARGUMENTS = 400
const CODE_ERROR_NO_SESSION = 401
const CODE_ERROR_NO_SIGNUP = 402
const CODE_ERROR_INTERNAL_SERVER_ERROR = 500

type Handler struct {
	us  *store.UserStore
	cfg *config.Config
}

func New(userStore *store.UserStore, config *config.Config) *Handler {
	return &Handler{
		us:  userStore,
		cfg: config,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	//Pages
	e.GET("/", h.DebugHandler)
	e.GET("/hello", h.helloTester)
	e.GET("/user/:id", h.userbyid)

	//Static
	e.File("/static/css", "static/css/style.css")
}
