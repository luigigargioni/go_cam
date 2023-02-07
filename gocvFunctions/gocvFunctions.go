package gocvFunctions

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gocv.io/x/gocv"
)

func GocvStartRecord(c echo.Context) error {
	fmt.Print("GocvStartRecord started\n")

	device := "0"
	if c.QueryParam("device") != "" {
		device = c.QueryParam("device")
	}

	prefix := ""
	if c.QueryParam("prefix") != "" {
		prefix = c.QueryParam("prefix") + "_"
	}

	format := ".mp4"
	if c.QueryParam("format") != "" {
		format = c.QueryParam("format")
		if !strings.HasPrefix(format, ".") {
			format = "." + format
		}
	}
	codec := "mp4v"
	switch format {
	case ".mp4":
		codec = "mp4v"
	case ".avi":
		codec = "XVID"
	}

	fileName := "videos/" + prefix + time.Now().Format("20060102150405") + format

	path := ""
	if c.QueryParam("path") != "" {
		path = c.QueryParam("path")
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
	}

	webcam, err := gocv.OpenVideoCapture(device)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", device)
		return c.String(http.StatusInternalServerError, "Error opening video capture device: "+device)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Cannot read device %v\n", device)
		return c.String(http.StatusInternalServerError, "Cannot read device: "+device)
	}

	writer, err := gocv.VideoWriterFile(path+fileName, codec, 25, img.Cols(), img.Rows(), true)
	if err != nil {
		fmt.Printf("Error opening video writer device: %v\n", fileName)
		return c.String(http.StatusInternalServerError, "Error opening video writer device: "+fileName)
	}
	defer writer.Close()

	for /* i := 0; i < 999999; i++ */ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", device)
			// return c.String(http.StatusInternalServerError, "Device closed: "+device)
			return c.String(http.StatusOK, fileName)
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

	device := "0"
	if c.QueryParam("device") != "" {
		device = c.QueryParam("device")
	}

	webcam, err := gocv.OpenVideoCapture(device)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", device)
		return c.String(http.StatusInternalServerError, "Error opening video capture device: "+device)
	}
	defer webcam.Close()

	return c.String(http.StatusOK, "GocvStopRecord success")
}
