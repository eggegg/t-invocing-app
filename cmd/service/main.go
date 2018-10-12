package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"


	"github.com/eggegg/t-invocing-app/handlers"
	"github.com/eggegg/t-invocing-app/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main()  {
	// create a new echo instance
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	  AllowOrigins: []string{"*"},
	  AllowHeaders: []string{"*"},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Signing Key for our auth middleware
	var signingKey = []byte("superdupersecret!")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.SigningContextKey, signingKey)
			return next(c)
		}
	})

	// add database to context
	db, err := sql.Open("mysql", "homestead:secret@tcp(localhost:33060)/t_invocing_app?charset=utf8") //第一个参数为驱动名  
	if err != nil {
		log.Fatalf("error opening database: %v\n", err)
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.DBContextKey, db)
			return next(c)
		}
	})


	// reminder handler group
	reminderGroup := e.Group("/invoice")
	reminderGroup.Use(middleware.JWT(signingKey))


	reminderGroup.GET("/hello", handlers.HelloWorld)

	reminderGroup.POST("", handlers.CreateInvoice) // create new invoice
	reminderGroup.GET("/user/:user_id", handlers.GetUserInvoice) // to fetch all invoices for a user
	reminderGroup.GET("/user/:user_id/:invoce_id", handlers.GetOneInvoice) // to fetch a certain invoice
	reminderGroup.POST("/sendmail", handlers.SendInvoice) // send a invoice to client

	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)

	// Authentication routes
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":8080"))
}