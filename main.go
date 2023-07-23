package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"studylesson/server"
)

func main() {

	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/tasks", server.List)
	e.GET("/tasks/:id", server.Get)
	s := http.Server{
		Addr:    ":8020",
		Handler: e,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
