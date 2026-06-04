package router

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"letter.go/arquive"
	"letter.go/brand"
	"letter.go/grammar"
	"letter.go/logic"
)

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

type User struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Color Color  `json:"color"`
}

type Observe struct {
	Sender   User   `json:"sender"`
	Language string `json:"language"`
	Message  string `json:"message"`
}

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

type Download struct {
	File string `json:"download"`
}

const (
	FILE_DIR      = "./raspberry"
	FILE_DIR_BAR  = "./raspberry/"
	FILE_TEMP     = "./temp"
	FILE_TEMP_BAR = "./temp/"
	ENVIRONMENT   = "Development"
	LOCAL         = "Production"
)

/*
func Controller(arbor grammar.Arbor, dome brand.Arbor, webcam *gocv.VideoCapture) *http.ServeMux {

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
		case "/Message":
			switch request.Method {
			case "POST":
				HandleMessage(writer, request, arbor)
			}
		case "/Translate":
			switch request.Method {
			case "POST":
				HandleTranslate(writer, request, arbor)
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
		case "/Model":
			switch request.Method {
			case "GET":
				HandleModel(writer, dome)
			}
		case "/Sentence":
			switch request.Method {
			case "GET":
				HandleSentence(writer, dome)
			}
		case "/Auxiliary":
			switch request.Method {
			case "GET":
				HandleAuxiliary(writer, dome)
			}
		case "/File":
			switch request.Method {
			case "POST":
				HandleFile(writer, request)
			case "DELETE":
				HandleDelete(writer, request)
			case "GET":
				HandleDownloadGet(writer, request)
			}
		case "/Download":
			switch request.Method {
			case "POST":
				HandleDownload(writer, request)
			case "GET":
				HandleDownloadGet(writer, request)
			}
		case "/Syntax":
			switch request.Method {
			case "POST":
				HandleSyntax(writer, request, arbor, dome)
			}
		case "/Stream":
			switch request.Method {
			case "GET":
				HandleStream(writer, request, webcam)
			}
		default:
			http.NotFound(writer, request)
		}
	})

	return mux
}
*/

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
		case "/Message":
			switch request.Method {
			case "POST":
				HandleMessage(writer, request, arbor)
			}
		case "/Translate":
			switch request.Method {
			case "POST":
				HandleTranslate(writer, request, arbor)
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
		case "/Model":
			switch request.Method {
			case "GET":
				HandleModel(writer, dome)
			}
		case "/Sentence":
			switch request.Method {
			case "GET":
				HandleSentence(writer, dome)
			}
		case "/Auxiliary":
			switch request.Method {
			case "GET":
				HandleAuxiliary(writer, dome)
			}
		case "/File":
			switch request.Method {
			case "POST":
				HandleFile(writer, request)
			case "DELETE":
				HandleDelete(writer, request)
			case "GET":
				HandleDownloadGet(writer, request)
			}
		case "/Download":
			switch request.Method {
			case "POST":
				HandleDownload(writer, request)
			case "GET":
				HandleDownloadGet(writer, request)
			}
		case "/Syntax":
			switch request.Method {
			case "POST":
				HandleSyntax(writer, request, arbor, dome)
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

func HandleMessage(writer http.ResponseWriter, request *http.Request, arbor grammar.Arbor) {
	defer request.Body.Close()

	var result Observe
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var language string = result.Language
	var phrase []grammar.Phrase = grammar.Snap(result.Message, arbor, language)
	response := phrase

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleTranslate(writer http.ResponseWriter, request *http.Request, arbor grammar.Arbor) {
	defer request.Body.Close()

	var result Observe
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var language string = result.Language
	var interpret []grammar.Interpret = grammar.Translate(result.Message, arbor, language)
	response := interpret

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

func HandleModel(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Model

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleSentence(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Sentence

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleAuxiliary(writer http.ResponseWriter, dome brand.Arbor) {
	value := dome.Auxiliary

	response := value

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

func HandleFile(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	file, header, err := request.FormFile("fileUpload")
	checkErr(err)
	defer file.Close()

	tempFile, err := os.CreateTemp("", "uploaded-*.txt")
	checkErr(err)

	_, err = io.Copy(tempFile, file)
	checkErr(err)

	fileContent, err := os.ReadFile(tempFile.Name())
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	var expression []byte = fileContent
	var register []int32 = []rune(header.Filename)

	var path_register string = ""
	if LOCAL == ENVIRONMENT {
		path_register = FILE_DIR_BAR + string(register)
	} else {
		path_register = string(register)
	}

	/*
		if len(name_file) > 0 {
				name_file = name_file[:len(name_file)]
			}
	*/

	arquive.Write(expression, path_register)
	checkErr(err)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write([]byte("Successful"))
}

func HandleDownload(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var result Download
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var file_name string = result.File
	var file_path string = ""
	if LOCAL == ENVIRONMENT {
		file_path = FILE_DIR_BAR + file_name
	} else {
		file_path = file_name
	}

	if len(file_name) == 0 {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Write([]byte("Unsuccessful"))
		return
	}
	file, err := os.Open(file_path)
	checkErr(err)
	defer file.Close()

	writer.Header().Set("Content-Disposition", "attachment; filename="+file_name)
	writer.Header().Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(writer, file)
	checkErr(err)
}

func HandleDownloadGet(writer http.ResponseWriter, request *http.Request) {
	var query = request.URL.Query()
	var name string = query.Get("name")

	var file_name string = name
	var file_path string = ""
	if LOCAL == ENVIRONMENT {
		file_path = FILE_DIR_BAR + file_name
	} else {
		file_path = file_name
	}

	if len(file_name) == 0 {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Write([]byte("Unsuccessful"))
		return
	}
	file, err := os.Open(file_path)
	checkErr(err)
	defer file.Close()

	writer.Header().Set("Content-Disposition", "attachment; filename="+file_name)
	writer.Header().Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(writer, file)
	checkErr(err)
}

func HandleDelete(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var result Download
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var file_name string = result.File
	var file_path string = ""
	if LOCAL == ENVIRONMENT {
		file_path = FILE_DIR_BAR + file_name
	} else {
		file_path = file_name
	}

	if len(file_name) == 0 {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Write([]byte("Unsuccessful"))
		return
	}

	defer os.Remove(file_path)
	checkErr(err)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write([]byte("Successful"))
}

func HandleSyntax(writer http.ResponseWriter, request *http.Request, arbor grammar.Arbor, dome brand.Arbor) {
	defer request.Body.Close()

	var result Observe
	var err = json.NewDecoder(request.Body).Decode(&result)
	checkErr(err)

	var language string = result.Language

	var phrase []grammar.Phrase = grammar.Snap(result.Message, arbor, language)
	var syntax []grammar.Recite = grammar.Syntax(phrase, dome, language)

	response := syntax

	responseJSON, err := json.Marshal(response)
	checkErr(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(responseJSON))
}

/*
func HandleStream(writer http.ResponseWriter, request *http.Request, webcam *gocv.VideoCapture) {
	writer.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	if !webcam.IsOpened() {
		return
	}
	var image gocv.Mat = gocv.NewMat()
	defer image.Close()
	for {
		var buffer *gocv.NativeByteBuffer
		buffer = stream.Read(webcam, image)
		writer.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
		writer.Write(buffer.GetBytes())
		writer.(http.Flusher).Flush()
	}
}
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
