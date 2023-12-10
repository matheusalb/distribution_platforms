package clientproxy

import (
	aux "Middleware/Auxiliar"
	"Middleware/Distribution/requestor"
	"reflect"
)

type ClientProxy struct {
	TypeName  string
	Host      string
	Port      int
	Id        int
	requestor requestor.Requestor
}

type ClientProxyBookSystem struct {
	Proxy ClientProxy
}

func NewClientProxy(host string, port int, id int) ClientProxy {
	return ClientProxy{Host: host, Port: port, Id: id, requestor: requestor.NewRequestor()}
}

func NewClientProxyBookSystem(host string, port int, id int) ClientProxyBookSystem {
	typeName := reflect.TypeOf(ClientProxyBookSystem{}).String()
	return ClientProxyBookSystem{
		ClientProxy{TypeName: typeName, Host: host, Port: port, Id: id}}
}

func (cp *ClientProxyBookSystem) DownloadBook(bookName string) string {
	params := make([]interface{}, 1)
	params[0] = bookName

	request := aux.Request{Op: "DownloadBook", Params: params}

	inv := aux.Invocation{Host: cp.Proxy.Host, Port: cp.Proxy.Port, Request: request}

	r := cp.Proxy.requestor.Invoke(inv)
	return r[0].(string)
}

func (cp *ClientProxy) Close() {
	request := aux.Request{Op: "Close", Params: nil}

	inv := aux.Invocation{Host: cp.Host, Port: cp.Port, Request: request}

	cp.requestor.Close(inv)
}
