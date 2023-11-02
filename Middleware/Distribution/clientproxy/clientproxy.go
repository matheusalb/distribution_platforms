package clientproxy

import (
	"Middleware/Distribution/requestor"
	shared "Middleware/Shared"
	"bufio"
	"fmt"
	"os"
)

type ClientProxy struct {
	Host string
	Port int
	Id   int
}

func NewClientProxy(host string, port int, id int) ClientProxy {
	return ClientProxy{Host: host, Port: port, Id: id}
}

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

func (cp *ClientProxy) DownloadBook(bookName string) int {
	params := make([]interface{}, 1)
	params[0] = bookName

	request := shared.Request{Op: "Download", Params: params}

	inv := shared.Invocation{Host: cp.Host, Port: cp.Port, Request: request}

	req := requestor.Requestor{}
	r := req.Invoke(inv)

	writeBook(bookName, r[0].(string))

	return 0
}
