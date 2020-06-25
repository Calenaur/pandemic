package handler

import (
	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	us  *store.UserStore
	ms  *store.MedicationStore
	cfg *config.Config
}

func New(userStore *store.UserStore, medicationStore *store.MedicationStore, config *config.Config) *Handler {
	return &Handler{
		us:  userStore,
		ms:  medicationStore,
		cfg: config,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	key := h.cfg.Token.Key
	//Pages
	//e.GET("/", h.DebugHandler)
	e.POST("/hello", h.helloTester)
	e.POST("/login", h.loginHandler)
	e.POST("/signup", h.signupHandler)
	//e.POST("/user/changename", h.changename)

	//Static
	e.File("/static/css", "static/css/style.css")

	//Restricted access, only for admin !TODO
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(key)))
	r.Use(middleware.CORS())
	r.GET("", restricted)
	r.GET("/user/:id", h.userbyid)
	r.GET("/users", h.listAll)
	r.PUT("/makeuseradmin", h.makeUserAdminHandler)
	r.DELETE("/user/:id", h.deleteUserByidHandler)

	//User specific stuff
	u := e.Group("/user")
	u.Use(middleware.JWT([]byte(key)))
	u.PUT("/username", h.changeNameHandler)
	u.PUT("/password", h.changePasswordHandler)
	u.DELETE("", h.deleteAccountHandler)
	u.GET("", h.getUserDetailsHandler)
	u.PUT("/balance", h.updateBalanceHandler)
	u.PUT("/manufacture", h.updateManufacture)
	u.GET("/diseases", h.getDiseasesHandler)
	u.GET("/available_diseases", h.getAvailableDiseasesHandler)
	u.PUT("/research_medication", h.medicationResearchHandler)
	u.GET("/friends", h.getFriendsHandler)
	u.POST("/friend_request", h.sendFriendRequestHandler)
	u.PUT("/friend_response", h.responseFriendRequestHandler)
	u.DELETE("/friend", h.deleteFriendHandler)
	//u.GET("/diseases_cures", h.whitchMedicationCuresWhichDiseaseHandler)

	//User Medication
	u.GET("/medication", h.getUserMedicationsHandler)
	u.GET("/medication/:id", h.getUserMedicationByIDHandler)

	//Medication
	m := e.Group("/medication")
	m.GET("", h.getMedicationsHandler)
	m.GET("/:id", h.getMedicationByIDHandler)
	m.GET("/trait", h.getMedicationTraitsHandler)
	m.GET("/trait/:id", h.getMedicationTraitByIDHandler)

}
