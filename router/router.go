package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	logic "letter.go/Logic"
	"letter.go/brand"
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

type Grammar struct {
	Kind string `json:"kind"`
}

func Controller(arbor grammar.Arbor, dome brand.Arbor) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			switch request.Method {
			case "GET":
				writer.Header().Set("Access-Control-Allow-Origin", "*")
				writer.Write([]byte("Grammar structure program."))
			case "POST":
				HandlePost(writer, request)
			}
		case "/Go":
			switch request.Method {
			case "POST":
				HandleGo(writer, request, arbor)
			}
		case "/Logic":
			switch request.Method {
			case "POST":
				HandleLogic(writer, request, arbor)
			}
		case "/Adverb":
			switch request.Method {
			case "GET":
				HandleAdverb(writer, dome)
			}
		case "/Pronoun":
			switch request.Method {
			case "GET":
				HandlePronoun(writer, dome)
			}
		case "/Article":
			switch request.Method {
			case "GET":
				HandleArticle(writer, dome)
			}
		case "/Conjunction":
			switch request.Method {
			case "GET":
				HandleConjunction(writer, dome)
			}
		case "/Numeral":
			switch request.Method {
			case "GET":
				HandleNumeral(writer, dome)
			}
		case "/Preposition":
			switch request.Method {
			case "GET":
				HandlePreposition(writer, dome)
			}
		case "/Verb":
			switch request.Method {
			case "GET":
				HandleVerb(writer, dome)
			}
		case "/Adjective":
			switch request.Method {
			case "GET":
				HandleAdjective(writer, dome)
			}
		case "/Noun":
			switch request.Method {
			case "GET":
				HandleNoun(writer, dome)
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

func HandleLogic(writer http.ResponseWriter, request *http.Request, arbor grammar.Arbor) {
	defer request.Body.Close()

	var result Notice
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var language string = "english"
	var phrase grammar.Phrase = grammar.Split(result.Message, arbor, language)
	var sense logic.Sense = logic.Math(phrase, result.Message)
	response := sense

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleAdverb(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Adverb

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandlePronoun(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Pronoun

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleArticle(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Article

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleConjunction(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Conjunction

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleNumeral(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Numeral

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandlePreposition(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Preposition

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleVerb(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Verb

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleAdjective(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Adjective

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleNoun(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Noun

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
