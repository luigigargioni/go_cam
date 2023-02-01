package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vee2xx/camtron"
)

func main() {
	e := echo.New()
	e.GET("/", testPath)
	e.GET("/start_record", startRecord)
	e.GET("/stop_record", stopRecord)
	e.Logger.Fatal(e.Start(":1323"))
}

func testPath(c echo.Context) error {
	return c.String(http.StatusOK, "It works!")
}

func startRecord(c echo.Context) error {
	fmt.Print("startRecord\n")
	camtron.StartStreamToFileConsumer() //start a listener that accepts and processes the stream
	camtron.StartCam()                  //start the Electron app that connects to the webcam and captures the stream
	camtron.StopWebcamUI()
	camtron.StartRecording()
	camtron.ShutdownStream()
	return c.String(http.StatusOK, "startRecord success")
}

func stopRecord(c echo.Context) error {
	fmt.Print("stopRecord\n")
	camtron.StopRecording()
	return c.String(http.StatusOK, "stopRecord success")
}
