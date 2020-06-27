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
	es  *store.EventStore
}

func New(userStore *store.UserStore, medicationStore *store.MedicationStore,
	eventStore *store.EventStore, config *config.Config) *Handler {
	return &Handler{
		us:  userStore,
		ms:  medicationStore,
		es:  eventStore,
		cfg: config,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	key := h.cfg.Token.Key
	//Pages
	//e.GET("/", h.DebugHandler)
	e.Use(middleware.CORS())
	e.POST("/hello", h.helloTester)
	e.POST("/login", h.loginHandler)
	e.POST("/signup", h.signupHandler)
	//e.POST("/user/changename", h.changename)

	//Static
	e.File("/static/css", "static/css/style.css")

	//Restricted access, only for admin !TODO
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(key)))
	r.GET("", restricted)
	r.GET("/user/:id", h.userbyid)
	r.GET("/users", h.listAll)
	r.GET("/usercount", h.userCount)
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

	//User Friend
	u.GET("/friend", h.getFriendsHandler)
	u.POST("/friend", h.sendFriendRequestHandler)
	u.PUT("/friend", h.responseFriendRequestHandler)
	u.DELETE("/friend", h.deleteFriendHandler)
	u.POST("/friend/gift", h.giftFriendHandler)

	//u.PUT("/change_tier", h.changeTierHandler)
	//u.GET("/diseases_cures", h.whitchMedicationCuresWhichDiseaseHandler)

	//User Event
	u.GET("/event", h.getEventsHandler)
	u.GET("/event/:id", h.getEventByIDHandler)
	u.GET("/event/mine", h.getMyEventsHandler)
	u.PUT("/event", h.subscribeToEventHandler)
	u.DELETE("/event", h.unSubscribeToEventHandler)

	//User Medication
	u.GET("/medication", h.getUserMedicationsHandler)
	u.GET("/medication/:id", h.getUserMedicationByIDHandler)

	//Medication
	m := e.Group("/medication")
	m.Use(middleware.JWT([]byte(key)))
	m.GET("", h.getMedicationsHandler)
	m.GET("/:id", h.getMedicationByIDHandler)
	m.GET("/trait", h.getMedicationTraitsHandler)
	m.GET("/trait/:id", h.getMedicationTraitByIDHandler)

}
