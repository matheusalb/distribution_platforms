package main

import (
	"fmt"
	"Middleware/Infrastructure/serverrequesthandlertcp"
	"io/ioutil"
	"os"
)

func lerLivro(nomeLivro string) []byte {
	file, err := os.Open("./books/"+nomeLivro + ".txt")
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
	srh := serverrequesthandlertcp.NewServerRequestHandlerTCP("localhost", 1313)
	for {
		nomeLivro := string(srh.Receive())
		livro := lerLivro(nomeLivro)

		srh.Send(livro)
	}
}