package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
	"letter.go/brand"
	"letter.go/grammar"
	"letter.go/router"
	"letter.go/sqlite"
)

func main() {
	var start = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("Start: ", start)
	var arbor grammar.Arbor = sqlite.Build()
	var dome brand.Arbor = sqlite.Forge()
	//var webcam *gocv.VideoCapture = stream.Video()
	var init = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("Init: ", init)
	//handler := cors.AllowAll().Handler(router.Controller(arbor, dome, webcam))
	handler := cors.AllowAll().Handler(router.Controller(arbor, dome))
	http.ListenAndServe(":8885", handler)
}
