package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"
	"letter.go/brand"
	"letter.go/grammar"
	"letter.go/router"
	"letter.go/sqlite"
)

func main() {
	var start time.Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("Start: ", start)
	var arbor grammar.Arbor = sqlite.Build()
	var dome brand.Arbor = sqlite.Forge()
	//var webcam *gocv.VideoCapture = flood.Video()
	var init time.Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("Init: ", init)

	//var handler http.Handler = cors.AllowAll().Handler(router.Controller(arbor, dome, webcam))
	handler := cors.AllowAll().Handler(router.Controller(arbor, dome))
	//http.ListenAndServe(":8885", handler)

	var failure *log.Logger = log.New(filtered{}, "", 0)

	var server *http.Server = &http.Server{
		Addr:         ":8885",
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		ErrorLog:     failure,
	}
	log.Fatal(server.ListenAndServe())
}

type filtered struct{}

func (out filtered) Write(result []byte) (number int, err error) {
	var message string = string(result)
	if strings.Contains(message, "wsasend") ||
		strings.Contains(message, "forcibly closed") ||
		strings.Contains(message, "broken pipe") ||
		strings.Contains(message, "connection reset by peer") ||
		strings.Contains(message, "i/o timeout") {
		return len(result), nil
	}
	return os.Stdout.Write(result)
}

/*
func isConnReset(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		var sysErr *os.SyscallError
		if errors.As(opErr.Err, &sysErr) {
			// Em sistemas Windows, o código é o WSAECONNRESET
			if errno, ok := sysErr.Err.(syscall.Errno); ok {
				const WSAECONNRESET = 10054
				return errno == syscall.ECONNRESET || errno == WSAECONNRESET
			}
		}
	}
	return false
}
*/
