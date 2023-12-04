package repository

import (
	"Middleware/Distribution/clientproxy"
	"calculadora/mymiddleware/client/proxies"
	"reflect"
)

func (proxy clientproxy.ClientProxy) interface{} {
	var clientProxy interface{}

	switch proxy.TypeName {
	case reflect.TypeOf(proxies.CalculatorProxy{}).String():
		calculatorProxy := proxies.NewCalculatorProxy()
		calculatorProxy.Proxy.TypeName = proxy.TypeName
		calculatorProxy.Proxy.Host = proxy.Host
		calculatorProxy.Proxy.Port = proxy.Port
		clientProxy = calculatorProxy
	}

	return clientProxy
}
