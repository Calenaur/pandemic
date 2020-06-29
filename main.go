package main

import (
	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/db"
	"github.com/calenaur/pandemic/handler"
	"github.com/calenaur/pandemic/store"
	echotemplate "github.com/foolin/echo-template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//Load config
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}

	//Connect to database
	con, err := db.New(cfg.Database.Username, cfg.Database.Password, cfg.Database.Database)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	//Setup stores
	userStore := store.NewUserStore(con, cfg)
	medicationStore := store.NewMedicationStore(con, cfg)
	userdataStore := store.NewUserdataStore(con, cfg)
  eventStore := store.NewEventStore(con, cfg)
	diseaseStore := store.NewDiseaseStore(con, cfg)

	//Setup handler
	handler := handler.New(userStore, medicationStore, userdataStore, eventStore, diseaseStore, cfg)


	//Setup echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = echotemplate.Default()

	//Routes
	handler.RegisterRoutes(e)
	port := cfg.Server.Port
	e.Logger.Fatal(e.Start(":" + port))
}
