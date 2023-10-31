package main

import (
	"Middleware/Infrastructure/clientrequesthandlertcp"
	"bufio"
	"fmt"
	"os"
)

func salvaLivro(titulo string, livro []byte) {
	file, err := os.Create("./books/" + titulo + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(livro)
	if err != nil {
		fmt.Println(err)
	}
	writer.Flush() // Não esqueça de chamar Flush para garantir que todos os dados sejam escritos no arquivo

	fmt.Println("Livro salvo com sucesso!")
}

func main() {
	var nomeLivro string
	crh := clientrequesthandlertcp.NewClientRequestHandlerTCP("localhost", 1313)

	fmt.Println("Digite o nome do Livro Desejado: ")
	fmt.Scanln(&nomeLivro)

	crh.Send([]byte(nomeLivro))
	livro := crh.Receive()

	salvaLivro(nomeLivro, livro)
}
