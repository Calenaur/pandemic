package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/foolin/echo-template"
	"github.com/calenaur/pandemic/db"
	"github.com/calenaur/pandemic/store"
	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/handler"
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

	//Setup handler
	handler := handler.New(userStore, cfg)

	//Setup echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = echotemplate.Default()

	//Routes
	handler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}