package main

import (
	"Middleware/Distribution/clientproxy"
	"Middleware/Distribution/invoker"
	"Middleware/Services/naming/proxy"
	shared "Middleware/Shared"
	"fmt"
)

func main() {

	fmt.Println("Running Server...")
	namingProxy := proxy.NamingProxy{}

	// Seleciona Porta Disponível
	port := shared.FindNextAvailablePort()

	clientProxyBookSystem := clientproxy.NewClientProxyBookSystem("localhost", port, 1)
	// Registro no Serviço de Nomes
	namingProxy.Register("BookSystem", clientProxyBookSystem)

	invoker := invoker.Invoker{}

	invoker.Invoke()
}
