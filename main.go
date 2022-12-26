package main

import (
	"github.com/folafunmi-db/go-echo-test/cmd/api/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/health", handler.HealthCheckHandler)
	e.GET("/posts", handler.PostsHandler)
	e.GET("/post/:id", handler.SinglePostHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
