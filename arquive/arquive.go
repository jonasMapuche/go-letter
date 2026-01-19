package arquive

import (
	"embed"
	"io/ioutil"
	"strings"
)

func Write(expression []byte, file string) {
	output(expression, file)
}

func output(expression []byte, file string) {
	var document []byte = expression
	var err_file = ioutil.WriteFile(file, document, 0644)
	if err_file != nil {
		panic(err_file)
	}
}

var diretory = "input"

//go:embed input*
var file embed.FS

func Input(arquive string) string {
	var file_path string = "input/" + arquive
	content, err := file.ReadFile(file_path)
	checkErr(err)
	var sql = string(content)
	sql = strings.ReplaceAll(sql, "\r", " ")
	sql = strings.ReplaceAll(sql, "\n", " ")
	return sql
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
