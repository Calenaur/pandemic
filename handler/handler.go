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
	es  *store.EventStore
	ds  *store.DiseaseStore
	cfg *config.Config
}

func New(userStore *store.UserStore, medicationStore *store.MedicationStore,
	eventStore *store.EventStore, diseaseStore *store.DiseaseStore, config *config.Config) *Handler {
	return &Handler{
		us:  userStore,
		ms:  medicationStore,
		es:  eventStore,
		ds:  diseaseStore,
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
	u.GET("/diseases", h.getDiseasesForUserHandler)
	u.GET("/available_diseases", h.getAvailableDiseasesHandler)
	u.PUT("/research_medication", h.medicationResearchHandler)
	u.GET("/friends", h.getFriendsHandler)
	u.POST("/friend_request", h.sendFriendRequestHandler)
	u.PUT("/friend_response", h.responseFriendRequestHandler)
	u.DELETE("/friend", h.deleteFriendHandler)
	u.PUT("/gift_friend", h.giftFriendHandler)
	//u.PUT("/change_tier", h.changeTierHandler)
	//u.GET("/diseases_cures", h.whitchMedicationCuresWhichDiseaseHandler)

	//User Event
	u.GET("/event", h.getEventsHandler)
	u.GET("/event/:id", h.getEventByIDHandler)

	//User Disease
	u.GET("/disease", h.getDiseasesHandler)
	u.GET("/disease/:id", h.getDiseaseByIDHandler)

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
