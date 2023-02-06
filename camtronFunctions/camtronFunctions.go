package camtronFunctions

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vee2xx/camtron"
)

func CamtronStartRecord(c echo.Context) error {
	fmt.Print("StartRecordCamtron started\n")
	camtron.StartStreamToFileConsumer() //start a listener that accepts and processes the stream
	camtron.StartCam()                  //start the Electron app that connects to the webcam and captures the stream
	// camtron.StopWebcamUI()
	// camtron.StartRecording()
	// camtron.ShutdownStream()
	return c.String(http.StatusOK, "StartRecordCamtron success")
}

func CamtronStopRecord(c echo.Context) error {
	fmt.Print("CamtronStopRecord started\n")
	camtron.StopRecording()
	return c.String(http.StatusOK, "CamtronStopRecord success")
}
