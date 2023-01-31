package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vee2xx/camtron"
)

func main() {
	e := echo.New()
	e.GET("/start_record", startRecord)
	e.GET("/stop_record", stopRecord)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.POST("/start_record", startRecord)
func startRecord(c echo.Context) error {
	camtron.StartStreamToFileConsumer()
	camtron.StartCam()
	return c.String(http.StatusOK, "startRecord")
}

// e.POST("/stop_record", stopRecord)
func stopRecord(c echo.Context) error {
	camtron.StopRecording()
	return c.String(http.StatusOK, "stopRecord")
}
