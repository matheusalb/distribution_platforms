package clientrequesthandlerudp

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type ClientRequestHandlerUDP struct {
	hostToConn string
	portToConn int
	addr       *net.UDPAddr
	conn       *net.UDPConn
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func NewClientRequestHandlerUDP(hostToConn string, portToConn int) *ClientRequestHandlerUDP {
	crh := new(ClientRequestHandlerUDP)
	crh.hostToConn = hostToConn
	crh.portToConn = portToConn
	crh.addr = nil
	crh.conn = nil

	return crh
}

func (crh *ClientRequestHandlerUDP) ConnectionUDP() {
	// Obtém o endereço do servidor para enviar os pacotes
	var err error

	crh.addr, err = net.ResolveUDPAddr("udp", crh.hostToConn+":"+strconv.Itoa(crh.portToConn))
	checkError(err)

	// Cria o socket UDP
	crh.conn, err = net.DialUDP("udp", nil, crh.addr)
	checkError(err)
}

func (crh *ClientRequestHandlerUDP) Send(msg []byte) {
	// Cria o socket UDP
	crh.ConnectionUDP()

	// Envia a mensagem
	_, err := crh.conn.Write(msg)
	checkError(err)
}

func (crh *ClientRequestHandlerUDP) Receive() []byte {

	msgFromServer := make([]byte, 8192)
	n, _, err := crh.conn.ReadFromUDP(msgFromServer)

	checkError(err)

	return msgFromServer[:n]
}
