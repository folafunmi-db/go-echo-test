package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/health", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostsHandler)
	e.GET("/post/:id", handlers.SinglePostHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
