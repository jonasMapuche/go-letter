package file

import (
	"io/ioutil"
	"os"
)

const path_directory string = "/" // linux = '/'
const directory string = "output/"
const File string = "letter.go"
const file_file string = "file.go"
const file_router string = "router.go"

func Body() {
	letter()
	output()
	file()
	router()
}

func letter() {
	var expression string = "package main\n\nimport (\n	\"fmt\"\n	\"time\"\n\n	\"gene.go/trees\"\n)\n\n"
	expression = expression + "type Notice struct {\n	Spawn	string\n	Date	time.Time\n	Quantity	float32\n}\n\n"
	expression = expression + "func main() {\n"
	expression = expression + "fmt.Println(\" -------------------------------------------------\")\n"
	expression = expression + "fmt.Println(\"| Start...\")\n"
	expression = expression + "fmt.Println(\" -------------------------------------------------\")\n"
	expression = expression + "fmt.Println(\"| Data...\")\n"
	expression = expression + "fmt.Println(\"|-------------------------------------------------\")\n"
	expression = expression + "fmt.Println(\"| Tree...\")\n"
	expression = expression + "var article Notice\n"
	expression = expression + "article.Spawn = \"raiz\"\n"
	expression = expression + "article.Date = time.Date(2023, 1, 17, 0, 0, 0, 0, time.Local)\n"
	expression = expression + "article.Quantity = 4\n"
	expression = expression + "var tree *trees.Tree = trees.Add(trees.Notice(article))\n"
	expression = expression + "fmt.Println(\"| Tree: \", tree.String())\n"
	expression = expression + "fmt.Println(\"|-------------------------------------------------\")\n"
	expression = expression + "}"

	var text = /*[]byte=*/ []byte(expression)
	var err_gene = /*error=*/ ioutil.WriteFile(directory+File, text, 0644)
	if err_gene != nil {
		panic(err_gene)
	}

}

func file() {
	var home string = directory + "files"
	var err_dir_file = /*error=*/ os.Mkdir(home, 0777)
	if err_dir_file != nil {
		panic(err_dir_file)
	}

	var expression string = "package files\n\nimport (\n	\"fmt\"\n)\n\n"
	expression = expression + "func Body() {\n"
	expression = expression + "fmt.Println(\" -------------------------------------------------\")\n"

	expression = expression + "fmt.Println(\"|-------------------------------------------------\")\n"
	expression = expression + "}"

	var text = /*[]byte*/ []byte(expression)
	var err_file = /*error=*/ ioutil.WriteFile(home+"/"+file_file, text, 0644)
	if err_file != nil {
		panic(err_file)
	}
}

func output() {
	var home string = directory + "output"
	var err_dir_output = os.Mkdir(home, 0777)
	if err_dir_output != nil {
		panic(err_dir_output)
	}
}

func router() {
	var home string = directory + "router"
	var err_dir_router = os.Mkdir(home, 0777)
	if err_dir_router != nil {
		panic(err_dir_router)
	}

	var expression string = "package router\n\nimport (\n	\"fmt\"\n)\n\n"
	expression = expression + "func Controller() {\n"
	expression = expression + "fmt.Println(\" -------------------------------------------------\")\n"

	expression = expression + "fmt.Println(\"|-------------------------------------------------\")\n"
	expression = expression + "}"

	var text = /*[]byte*/ []byte(expression)
	var err_router = /*error=*/ ioutil.WriteFile(home+"/"+file_router, text, 0644)
	if err_router != nil {
		panic(err_router)
	}
}
