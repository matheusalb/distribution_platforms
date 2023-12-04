package proxy

import (
	aux "Middleware/Auxiliar"
	"Middleware/Distribution/clientproxy"
	"Middleware/Distribution/requestor"
	repository "Middleware/Repository"
	shared "Middleware/Shared"
	"fmt"
)

type NamingProxy struct{}

func (NamingProxy) Register(p1 string, proxy interface{}) bool {

	// prepare invocation
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = proxy
	namingproxy := clientproxy.ClientProxy{Host: "", Port: shared.N_PORT_NS, Id: 0}
	request := aux.Request{Op: "Register", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv) //.([]interface{})

	return ter[0].(bool)
}

func (NamingProxy) Lookup(p1 string) interface{} {
	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	namingproxy := clientproxy.ClientProxy{Host: "", Port: shared.N_PORT_NS, Id: 0}
	request := aux.Request{Op: "Lookup", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv) //.([]interface{})

	// process reply
	proxyTemp := ter[0].(map[string]interface{})
	fmt.Println(proxyTemp)
	fmt.Println("--------------------")
	clientProxyTemp := clientproxy.ClientProxy{TypeName: proxyTemp["TypeName"].(string), Host: proxyTemp["Host"].(string), Port: int(proxyTemp["Port"].(float64))}
	fmt.Println(clientProxyTemp)
	clientProxy := repository.CheckRepository(clientProxyTemp)

	return clientProxy
}
