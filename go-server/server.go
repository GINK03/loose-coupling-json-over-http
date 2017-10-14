package main

import (
	_ "fmt"
	"github.com/labstack/echo"
	"net/http"
)

type (
	Input struct {
		Type string  `json:"type" form:"type" query:"type"`
		Data float64 `json:"data" form:"data" query:"data"`
	}
)

func main() {
	e := echo.New()
	e.POST("/test", func(c echo.Context) error {
		u := new(Input)

		c.Bind(u)
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
