package shared

import (
	"log"
)

const MIOP_REQUEST = 1
const MIOP_REPLY = 2

const N_HOST_SERVIDOR = "localhost"
const N_PORT_SERVIDOR = 1515

type Request struct {
	Op     string
	Params []interface{}
}

type Invocation struct {
	Host    string
	Port    int
	Request Request
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}
