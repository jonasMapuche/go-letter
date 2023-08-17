package router

import (
	"io/ioutil"
	"net/http"
)

const URL = "http://localhost:8100"

func Controller() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			switch request.Method {
			case "GET":
				writer.Header().Set("Access-Control-Allow-Origin", "*")
				writer.Write([]byte(" {'message': 'Hello World!'} "))
			case "POST":
				HandlePost(writer, request)
			}
		default:
			http.NotFound(writer, request)
		}
	})

	return mux

}

func HandlePost(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.Write([]byte("We have a problem reading the content! Please, try again."))
		writer.WriteHeader(http.StatusBadGateway)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
}
