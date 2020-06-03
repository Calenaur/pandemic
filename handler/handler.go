package handler

import (
	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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
	//e.GET("/", h.DebugHandler)
	e.POST("/hello", h.helloTester)
	e.GET("/user/:id", h.userbyid)
	e.POST("/login", h.loginHandler)
	e.POST("/signup", h.signupHandler)

	//Static
	e.File("/static/css", "static/css/style.css")

	//Grouping
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("جامعة هانزه العلوم تطبيقية")))
	r.GET("", restricted)

}
