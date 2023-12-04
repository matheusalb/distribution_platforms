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

	clientproxy_tmp := namingProxy.Lookup("BookSystem")
	fmt.Println(clientproxy_tmp)

	clientProxy := clientproxy_tmp.(clientproxy.ClientProxyBookSystem)
	var nomeLivro string
	// clientproxy := clientproxy.NewClientProxy(shared.N_HOST_SERVIDOR, shared.N_PORT_SERVIDOR, 1)

	fmt.Println("Digite o nome do Livro Desejado: ")
	fmt.Scanln(&nomeLivro)
	book := clientProxy.DownloadBook(nomeLivro)
	writeBook(nomeLivro, book)
}
