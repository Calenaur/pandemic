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
	cfg *config.Config
}

func New(userStore *store.UserStore, medicationStore *store.MedicationStore, userdataStore *store.UserdataStore, config *config.Config) *Handler {
	return &Handler{
		us:  userStore,
		ms:  medicationStore,
		ud:  userdataStore,
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
	u.GET("/disease", h.getDiseasesHandler)
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

	//Userdata
	u.POST("/disease", h.setUserDiseaseHandler) //FormValue diseaseid
	// u.GET("/disease", h.getUserDiseaseHandler)
	u.POST("/tier", h.setUserTierHandler) //FormValue tier
	u.GET("/tier", h.getUserTierHandler)
	u.POST("/event", h.setUserEventHandler) // FormValue event
	u.GET("/event", h.getUserEventHandler)
	u.POST("/researcher", h.setUserResearcherHandler) // FormValue researcher, researchername
	u.GET("/researcher", h.getUserResearcherHandler)
	u.PUT("/researcher", h.updateUserResearcherHandler)         // FormValue oldname, newname
	u.POST("/researchertrait", h.setUserResearcherTraitHandler) // FormValue researchername, traitname
	u.GET("/researchertrait", h.getUserResearcherTraitHandler)

	//Medication
	m := e.Group("/medication")
	m.GET("", h.getMedicationsHandler)
	m.GET("/:id", h.getMedicationByIDHandler)
	m.GET("/trait", h.getMedicationTraitsHandler)
	m.GET("/trait/:id", h.getMedicationTraitByIDHandler)

}
