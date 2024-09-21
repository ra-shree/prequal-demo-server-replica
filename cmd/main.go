package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	middleware "github.com/ra-shree/prequal-demo-server-replica/pkg/middlewares"
)

func main() {
	e := echo.New()

	e.Use(middleware.Process)

	e.GET("/ping", middleware.HandleGetPing)

	e.GET("/test", func(c echo.Context) error {
		randomDelay := time.Duration(rand.Intn(500)+100) * time.Millisecond
		time.Sleep(randomDelay)
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1233"))
}
