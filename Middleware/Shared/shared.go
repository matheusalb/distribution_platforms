package shared

import (
	"log"
	"net"
	"strconv"
)

const MIOP_REQUEST = 1
const MIOP_REPLY = 2

const N_HOST_SERVIDOR = "localhost"
const N_PORT_SERVIDOR = 1515

const N_PORT_NS = 9090

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}

func FindNextAvailablePort() int { // TCP only
	i := 3000

	for i = 3000; i < 4000; i++ {
		port := strconv.Itoa(i)
		ln, err := net.Listen("tcp", ":"+port)

		if err == nil {
			ln.Close()
			break
		}
	}
	return i
}
