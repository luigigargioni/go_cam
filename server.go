package main

import (
	"go_cam/gocvFunctions"
	"net/http"

	"github.com/labstack/echo/v4"
)

const PORT = ":1323"

func main() {
	e := echo.New()
	e.GET("/", TestPath)

	// GoCV
	e.GET("/start_record", gocvFunctions.GocvStartRecord)
	e.GET("/stop_record", gocvFunctions.GocvStopRecord)

	e.Logger.Fatal(e.Start(PORT))
}

func TestPath(c echo.Context) error {
	return c.String(http.StatusOK, "It works!")
}
