package main

import (
	"Middleware/Infrastructure/serverrequesthandlerudp"
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
	srh := serverrequesthandlerudp.NewServerRequestHandlerUDP("localhost", 1313)
	for {
		nomeLivro := string(srh.Receive())
		livro := lerLivro(nomeLivro)

		srh.Send(livro)
	}
}
