package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"letter.go/grammar"
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

func Controller(arbor grammar.Arbor) *http.ServeMux {

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
		case "/Pronoun":
			switch request.Method {
			case "POST":
				HandlePronoun(writer, request, arbor)
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

func HandleGo(writer http.ResponseWriter, request *http.Request, arbor grammar.Arbor) {
	defer request.Body.Close()

	var result Notice
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var language string = "english"
	var phrase grammar.Phrase = grammar.Split(result.Message, arbor, language)
	response := phrase

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandlePronoun(writer http.ResponseWriter, request *http.Request, arbor grammar.Arbor) {
	defer request.Body.Close()
	var result Notice
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var language string = "english"
	var phrase grammar.Phrase = grammar.Split(result.Message, arbor, language)
	var value grammar.Phrase = grammar.Agree(phrase)

	response := value

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
