package arquive

import (
	"io/ioutil"
	"os"
	"strings"
)

const directory = "input/"

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

func Input(arquive string) string {
	var file_path string = directory + arquive
	content, err := os.ReadFile(file_path)
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
