package main

import (
	"Middleware/Distribution/clientproxy"
	shared "Middleware/Shared"
	"fmt"
)

func main() {
	var nomeLivro string
	clientproxy := clientproxy.NewClientProxy(shared.N_HOST_SERVIDOR, shared.N_PORT_SERVIDOR, 1)

	fmt.Println("Digite o nome do Livro Desejado: ")
	fmt.Scanln(&nomeLivro)
	clientproxy.DownloadBook(nomeLivro)
}
