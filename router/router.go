package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"letter.go/tree"
)

const URL = "http://localhost:8100"

type Notice struct {
	Message string `json:"message"`
}

type Answer struct {
	Message string `json:"message"`
	Class   string `json:"class"`
	Kind    string `json:"kind"`
}

func Controller(arbor tree.Arbor) *http.ServeMux {

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
		case "/Go":
			switch request.Method {
			case "POST":
				HandleGo(writer, request, arbor)
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

func HandleGo(writer http.ResponseWriter, request *http.Request, arbor tree.Arbor) {
	defer request.Body.Close()

	var result Notice
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var phrase tree.Phrase = tree.Split(result.Message, arbor)

	var message string = ""
	var class string = ""
	for _, value := range phrase.Word {
		if (message == "") || (class == "") {
			message = value.Term
			class = value.Class
		} else {
			message = message + ", " + value.Term
			class = class + ", " + value.Class
		}
	}

	response := Answer{
		Message: message,
		Class:   class,
		Kind:    phrase.Kind,
	}

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
