package main

import (
	"Middleware/Distribution/invoker"
	"fmt"
	"io/ioutil"
	"os"
)

func lerLivro(nomeLivro string) []byte {
	file, err := os.Open("./books/" + nomeLivro + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return content
}

func main() {
	invoker := invoker.Invoker{}

	invoker.Invoke()
}
