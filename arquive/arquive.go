package arquive

import (
	"io/ioutil"
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
