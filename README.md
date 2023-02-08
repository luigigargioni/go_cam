# :camera: GO_CAM

## :dart: Requirements

* SO: Windows

* GO Library: [Echo](https://echo.labstack.com/)

* HTTP Server

* Record without audio

* Endpoints:
  * startRecord GET ```start_record```
    * parameter ```device``` (default 0)
    * parameter ```path``` (default "")
    * parameter ```prefix``` (default "")
    * parameter ```format``` (default "mp4")
    * registration name: ```<client-prefix>_<timestamp>.<video-format>```
    * respond with generated name
  * stopRecord GET ```stop_record```
    * parameter ```device``` (default 0)

---

## :wrench: Start dev-server

```bash
go run server.go
```

For Hot Reload ([Air](https://github.com/cosmtrek/air)):

```bash
air
```

---

## :hammer: Implementation

### :one: [Camtron](https://github.com/vee2xx/camtron)

* Electron port: ```8080```

* Endpoints
  * start_record_camtron
  * stop_record_camtron

### :two: [GoCV](https://github.com/hybridgroup/gocv)

* [Installation video](https://www.youtube.com/watch?v=c2HbPpEFYIA)

* [mingw-w64 problems](https://stackoverflow.com/questions/46455927/mingw-w64-installer-the-file-has-been-downloaded-incorrectly)

* Endpoints
  * start_record_gocv
  * stop_record_gocv

### :three: Linux

Libraries:

* [Linux Video API V4L2](https://medium.com/learning-the-go-programming-language/realtime-video-capture-with-go-65a8ac3a57da)

* [go-webcam](https://github.com/blackjack/webcam) (base on V4L2)

* [go4vl](https://medium.com/go4vl/building-a-webcam-with-go-and-go4vl-7b56d2c54e39)

### :four: Frontend

* [Plain Javascript](https://web.dev/media-recording-video/)

* [vue-web-cam](https://www.npmjs.com/package/vue-web-cam)
