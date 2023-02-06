package gocvFunctions

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gocv.io/x/gocv"
)

func GocvStartRecord(c echo.Context) error {
	fmt.Print("GocvStartRecord started\n")

	deviceID := c.QueryParam("deviceID")
	saveFile := c.QueryParam("saveFile")

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return c.String(http.StatusInternalServerError, "Error opening video capture device: "+deviceID)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Cannot read device %v\n", deviceID)
		return c.String(http.StatusInternalServerError, "Cannot read device: "+deviceID)
	}

	writer, err := gocv.VideoWriterFile(saveFile, "mp4v", 25, img.Cols(), img.Rows(), true)
	if err != nil {
		fmt.Printf("Error opening video writer device: %v\n", saveFile)
		return c.String(http.StatusInternalServerError, "Error opening video writer device: "+saveFile)
	}
	defer writer.Close()

	for i := 0; i < 100; i++ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return c.String(http.StatusInternalServerError, "Device closed: "+deviceID)
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
	return c.String(http.StatusOK, "GocvStartRecord success")
}

func GocvStopRecord(c echo.Context) error {
	fmt.Print("GocvStopRecord started\n")
	return c.String(http.StatusOK, "GocvStopRecord success")
}
