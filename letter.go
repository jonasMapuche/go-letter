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
	"letter.go/tree"
)

func main() {
	var start = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println(" -------------------------------------------------")
	fmt.Println("| Start...")
	fmt.Println("| Date and time: ", start)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| SQLite...")
	fmt.Println("| Grammar...")
	var arbor grammar.Arbor = sqlite.Build()
	var end = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("| Date and time: ", end)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| SQLite...")
	fmt.Println("| Grammar...")
	var dome brand.Arbor = sqlite.Forge()
	end = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("| Date and time: ", end)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Tree...")
	tree.Plant()
	end = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("| Date and time: ", end)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| File...")
	//file.Body()
	//fmt.Println("| File:", file.File)
	end = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	fmt.Println("| Date and time: ", end)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Http...")
	handler := cors.AllowAll().Handler(router.Controller(arbor, dome))
	http.ListenAndServe(":8885", handler)
}
