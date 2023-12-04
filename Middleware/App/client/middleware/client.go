package main

import (
	"Middleware/Distribution/clientproxy"
	"Middleware/Services/naming/proxy"
	"bufio"
	"fmt"
	"os"
)

func writeBook(title string, book string) {
	file, err := os.Create("./books/" + title + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(book))
	if err != nil {
		fmt.Println(err)
	}
	writer.Flush()
}

func main() {
	fmt.Println("Running Client...")
	namingProxy := proxy.NamingProxy{}

	clientProxy := namingProxy.Lookup("BookSystem").(clientproxy.ClientProxyBookSystem)

	defer func(clientProxy *clientproxy.ClientProxy) {
		fmt.Println("Closing Client...")
		clientProxy.Close()
	}(&clientProxy.Proxy)

	var nomeLivro string

	for {
		fmt.Println("Digite o nome do Livro Desejado: ")
		fmt.Scanln(&nomeLivro)
		if nomeLivro == "exit" {
			break
		}

		book := clientProxy.DownloadBook(nomeLivro)
		fmt.Println(book[1:20])
	}
	// writeBook(nomeLivro, book)
}
