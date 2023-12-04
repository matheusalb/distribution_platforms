package main

import (
	"Middleware/Services/naming/invoker"
	"fmt"
)

func main() {

	fmt.Println("Naming servidor running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
