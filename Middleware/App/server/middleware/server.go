package main

import (
	"Middleware/Distribution/invoker"
)

func main() {
	invoker := invoker.Invoker{}

	invoker.Invoke()
}
