package handler

import (
	"net/http"
	"strconv"

	"github.com/folafunmi-db/go-echo-test/cmd/api/service"

	"github.com/labstack/echo"
)

func PostsHandler(c echo.Context) error {
	data, err := service.GetAll()

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process this request")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}

func SinglePostHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process this request")
	}

	data, err := service.GetById(idx)

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process this request")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}
