package main

import (
	"Middleware/Distribution/clientproxy"
	shared "Middleware/Shared"
	"bufio"
	"fmt"
	"os"
)

func writeBook(title string, book []byte) {
	file, err := os.Create("./books/" + title + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(book)
	if err != nil {
		fmt.Println(err)
	}
	writer.Flush()
}

func main() {
	var nomeLivro string
	clientproxy := clientproxy.NewClientProxy(shared.N_HOST_SERVIDOR, shared.N_PORT_SERVIDOR, 1)

	fmt.Println("Digite o nome do Livro Desejado: ")
	fmt.Scanln(&nomeLivro)
	book := clientproxy.DownloadBook(nomeLivro)
	writeBook(nomeLivro, book)
}
