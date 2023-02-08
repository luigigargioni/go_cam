# :camera: GO_CAM

## :dart: Requirements

* SO: Windows

* Go HTTP Server Library: [Echo](https://echo.labstack.com/)

* Go Vision Library: [GoCV](https://gocv.io/)

* Record without audio

* Endpoints:
  * startRecord &rarr; ```GET start_record```
    * parameter ```device```: Camera to use. Default ```0```
    * parameter ```path```: Path where to save video. Default ```""```
    * parameter ```prefix```: Prefix for video filename. Default ```""```
    * parameter ```format```: Video format. Default ```mp4```
    * registration name: ```<prefix>_<timestamp>.<format>```
    * respond with generated name: e.g. ```prefix_20230208145236.mp4```
  * stopRecord &rarr; ```GET stop_record```
    * parameter ```device```: Camera to stop. Default ```0```

---

## :wrench: Start dev-server

* Default Port: ```1323```

```bash
go run server.go
```

For Hot Reload ([Air](https://github.com/cosmtrek/air)):

```bash
air
```

---

## :hammer: Implementation

### [GoCV](https://gocv.io/)

* [GitHub](https://github.com/hybridgroup/gocv)

* [Windows Installing](https://gocv.io/getting-started/windows/)

* [Installation video](https://www.youtube.com/watch?v=c2HbPpEFYIA)

* [mingw-w64 problems with Windows](https://stackoverflow.com/questions/46455927/mingw-w64-installer-the-file-has-been-downloaded-incorrectly)

* [Linux Installing](https://gocv.io/getting-started/linux/)

---

## :arrows_clockwise: Other solutions

### :one: [Camtron](https://github.com/vee2xx/camtron)

* Electron port: ```8080```

### :two: Linux libraries

* [Linux Video API V4L2](https://medium.com/learning-the-go-programming-language/realtime-video-capture-with-go-65a8ac3a57da)

* [go-webcam](https://github.com/blackjack/webcam) (base on V4L2)

* [go4vl](https://medium.com/go4vl/building-a-webcam-with-go-and-go4vl-7b56d2c54e39)

### :three: Frontend Vue.js

* [Plain Javascript](https://web.dev/media-recording-video/)

* [vue-web-cam](https://www.npmjs.com/package/vue-web-cam)
