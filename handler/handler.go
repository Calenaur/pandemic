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
	key := h.cfg.Token.Key
	//Pages
	//e.GET("/", h.DebugHandler)
	e.POST("/hello", h.helloTester)
	e.GET("/usr/:id", h.userbyid)
	e.POST("/login", h.loginHandler)
	e.POST("/signup", h.signupHandler)
	e.GET("/users/:page", h.listAll)
	//e.POST("/user/changename", h.changename)

	//Static
	e.File("/static/css", "static/css/style.css")

	//Restricted access, only for admin !TODO
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(key)))
	e.Use(middleware.CORS())
	r.GET("", restricted)

	//User specific stuff
	u := e.Group("/user")
	u.Use(middleware.JWT([]byte(key)))
	u.PUT("/changename", h.changeNameHandler)
	u.PUT("/changepassword", h.changePasswordHandler)
	u.DELETE("/deleteaccount", h.deleteAccountHandler)
	u.GET("", h.getUserDetailsHandler)
	u.PUT("/balance", h.updateBalanceHandler)
	u.PUT("/manufacture", h.updateManufacture)

}
