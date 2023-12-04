package clientproxy

import (
	aux "Middleware/Auxiliar"
	"Middleware/Distribution/requestor"
	"fmt"
	"reflect"
)

type ClientProxy struct {
	TypeName string
	Host     string
	Port     int
	Id       int
}

type ClientProxyBookSystem struct {
	Proxy ClientProxy
}

func NewClientProxy(host string, port int, id int) ClientProxy {
	return ClientProxy{Host: host, Port: port, Id: id}
}

func NewClientProxyBookSystem(host string, port int, id int) ClientProxyBookSystem {
	typeName := reflect.TypeOf(ClientProxyBookSystem{}).String()
	fmt.Println(typeName)
	return ClientProxyBookSystem{
		ClientProxy{TypeName: typeName, Host: host, Port: port, Id: id}}
}

func (cp *ClientProxyBookSystem) DownloadBook(bookName string) string {
	params := make([]interface{}, 1)
	params[0] = bookName

	request := aux.Request{Op: "Download", Params: params}

	inv := aux.Invocation{Host: cp.Proxy.Host, Port: cp.Proxy.Port, Request: request}

	req := requestor.Requestor{}
	r := req.Invoke(inv)

	return r[0].(string)
}
