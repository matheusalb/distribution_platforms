package impl

import (
	"fmt"
	"io/ioutil"
	"os"
)

type BookService struct {
	Id int
}

func (service *BookService) DownloadBook(nomeLivro string) string {
	file, err := os.Open("./books/" + nomeLivro + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}
