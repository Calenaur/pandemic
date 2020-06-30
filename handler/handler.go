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
	ud  *store.UserdataStore
	es  *store.EventStore
	ds  *store.DiseaseStore
	cfg *config.Config
}

func New(userStore *store.UserStore, medicationStore *store.MedicationStore, userdataStore *store.UserdataStore, eventStore *store.EventStore, diseaseStore *store.DiseaseStore, config *config.Config) *Handler {
	return &Handler{
		us:  userStore,
		ms:  medicationStore,
		ud:  userdataStore,
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
	u.POST("/device", h.updateDeviceHandler)
	u.GET("/device", h.getDeviceHandler)

	//User Friend
	u.GET("/friend", h.getFriendsHandler)
	u.POST("/friend", h.sendFriendRequestHandler)
	u.PUT("/friend", h.responseFriendRequestHandler)
	u.DELETE("/friend", h.deleteFriendHandler)
	u.POST("/friend/gift", h.giftFriendHandler)
	u.GET("/friend/pending", h.pendingFriendsHandler)

	//u.PUT("/change_tier", h.changeTierHandler)

	//User Event
	u.GET("/event", h.getMyEventsHandler)
	u.PUT("/event", h.subscribeToEventHandler)
	u.DELETE("/event", h.unSubscribeToEventHandler)

	//User Disease
	u.GET("/disease", h.getDiseasesForUserHandler)
	u.GET("/disease/available", h.getAvailableDiseasesHandler)
	u.POST("/disease", h.setUserDiseaseHandler) //FormValue diseaseid
	u.DELETE("/disease", h.unSelectDiseaseHandler)

	//User Medication
	u.GET("/medication", h.getUserMedicationsHandler)
	u.GET("/medication/:id", h.getUserMedicationByIDHandler)
	u.PUT("/medication", h.medicationResearchHandler)

	//**Userdata**
	//User Tier
	u.POST("/tier", h.setUserTierHandler) //FormValue tier
	u.GET("/tier", h.getUserTierHandler)

	//User Researcher
	u.POST("/researcher", h.setUserResearcherHandler) // FormValue researcher, researchername
	u.GET("/researcher", h.getUserResearcherHandler)
	u.PUT("/researcher", h.updateUserResearcherHandler) // FormValue oldname, newname

	//User Researcher Trait
	u.POST("/researchertrait", h.setUserResearcherTraitHandler) // FormValue researchername, traitname
	u.GET("/researchertrait", h.getUserResearcherTraitHandler)

	//Medication
	m := e.Group("/medication")
	m.Use(middleware.JWT([]byte(key)))
	m.GET("", h.getMedicationsHandler)
	m.GET("/:id", h.getMedicationByIDHandler)
	m.GET("/trait", h.getMedicationTraitsHandler)
	m.GET("/trait/:id", h.getMedicationTraitByIDHandler)

	//Event
	ev := e.Group("/event")
	ev.Use(middleware.JWT([]byte(key)))
	ev.GET("", h.getEventsHandler)
	ev.GET("/:id", h.getEventByIDHandler)

	//Disease
	d := e.Group("/disease")
	d.Use(middleware.JWT([]byte(key)))
	d.GET("", h.getDiseasesHandler)
	d.GET("/:id", h.getDiseaseByIDHandler)
	d.GET("/medication", h.getDiseaseMedicationHandler)

	//Duplicates
	// u.POST("/disease", h.selectDiseaseHandler)
	// u.GET("/disease", h.getUserDiseaseHandler)
	// u.POST("/event", h.setUserEventHandler) // FormValue event
	// u.GET("/event", h.getUserEventHandler)
}
