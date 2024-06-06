package main

import (
	"github.com/emarifer/go-echo-templ-htmx/db"
	"github.com/emarifer/go-echo-templ-htmx/handlers"
	"github.com/emarifer/go-echo-templ-htmx/services"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// In production, the secret key of the CookieStore
// and database name would be obtained from a .env file
const (
	SECRET_KEY string = "secret"
	DB_NAME    string = "app_data.db"
)

func main() {

	e := echo.New()

	e.Static("/", "assets")

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// Helpers Middleware
	// e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Session Middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

	store, err := db.NewStore(DB_NAME)
	if err != nil {
		e.Logger.Fatalf("failed to create store: %s", err)
	}

	userServices := services.NewUserServices(services.User{}, store)
	authHandler := handlers.NewAuthHandler(userServices)

	todoServices := services.NewTodoServices(services.Todo{}, store)
	taskHandler := handlers.NewTaskHandler(todoServices)

	// Setting Routes
	handlers.SetupRoutes(e, authHandler, taskHandler)

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}

/*
https://gist.github.com/taforyou/544c60ffd072c9573971cf447c9fea44
https://gist.github.com/mhewedy/4e45e04186ed9d4e3c8c86e6acff0b17

https://github.com/CurtisVermeeren/gorilla-sessions-tutorial
*/
