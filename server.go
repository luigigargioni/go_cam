package main

import (
	"go_cam/camtronFunctions"
	"go_cam/gocvFunctions"
	"net/http"

	"github.com/labstack/echo/v4"
)

const PORT = ":1323"

func main() {
	e := echo.New()
	e.GET("/", TestPath)

	// Camtron
	e.GET("/start_record_camtron", camtronFunctions.CamtronStartRecord)
	e.GET("/stop_record_camtron", camtronFunctions.CamtronStopRecord)

	// GoCV
	e.GET("/start_record_gocv", gocvFunctions.GocvStartRecord)
	e.GET("/stop_record_gocv", gocvFunctions.GocvStopRecord)

	e.Logger.Fatal(e.Start(PORT))
}

func TestPath(c echo.Context) error {
	return c.String(http.StatusOK, "It works!")
}
