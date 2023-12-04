package repository

import (
	"Middleware/Distribution/clientproxy"
	"fmt"
	"reflect"
)

func CheckRepository(proxy clientproxy.ClientProxy) interface{} {
	var clientProxy interface{}

	fmt.Println(reflect.TypeOf(clientproxy.ClientProxyBookSystem{}).String())
	fmt.Println("--------------------")
	switch proxy.TypeName {
	case reflect.TypeOf(clientproxy.ClientProxyBookSystem{}).String():
		bookProxy := clientproxy.NewClientProxyBookSystem(proxy.Host, proxy.Port, proxy.Id)
		bookProxy.Proxy.TypeName = proxy.TypeName
		bookProxy.Proxy.Host = proxy.Host
		bookProxy.Proxy.Port = proxy.Port
		clientProxy = bookProxy
	}

	return clientProxy
}
