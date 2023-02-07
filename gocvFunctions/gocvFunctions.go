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
	codec := GetCodec(format)

	fileName := prefix + time.Now().Format("20060102150405") + format

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

	img := gocv.NewMat()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Cannot read device %v\n", device)
		return c.String(http.StatusInternalServerError, "Cannot read device: "+device)
	}

	FPS := 25.0
	WIDTH := img.Cols()
	HEIGHT := img.Rows()
	IS_COLOR := true
	writer, err := gocv.VideoWriterFile(path+fileName, codec, FPS, WIDTH, HEIGHT, IS_COLOR)
	if err != nil {
		fmt.Printf("Error opening video writer device: %v\n", fileName)
		return c.String(http.StatusInternalServerError, "Error opening video writer device: "+fileName)
	}

	// Start recording in a new goroutine
	go StartRecording(webcam, writer, img, device, fileName, FPS)

	return c.String(http.StatusOK, "GocvStartRecord success\nFilename: "+fileName)
}

func StartRecording(webcam *gocv.VideoCapture, writer *gocv.VideoWriter, img gocv.Mat, device string, fileName string, fps float64) {
	defer img.Close()
	defer writer.Close()
	MAX_SECONDS := 300 // 5 minutes
	for i := 0; i < MAX_SECONDS*int(fps); i++ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", device)
			return
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
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

func GetCodec(format string) string {
	codec := ""
	switch format {
	case ".mp4":
		codec = "mp4v"
	case ".avi":
		codec = "XVID"
	// Add more formats here
	default:
		codec = "mp4v"
	}

	return codec
}
