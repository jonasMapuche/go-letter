package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
	"letter.go/router"
)

func main() {
	var start = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println(" -------------------------------------------------")
	fmt.Println("| Start...")
	fmt.Println("| Date and time: ", start)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Data...")
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Tree...")
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| File...")
	//file.Body()
	//fmt.Println("| File:", file.File)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Http...")

	handler := cors.AllowAll().Handler(router.Controller())

	http.ListenAndServe(":8885", handler)
	fmt.Println("| End...")
	var end = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("| Date and time: ", end)
	fmt.Println(" -------------------------------------------------")
}
