package serverrequesthandlerudp

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type ServerRequestHandlerUDP struct {
	serverHost string
	serverPort int
	conn       *net.UDPConn // Socket UDP
	addr       *net.UDPAddr // Endereço UDP do cliente
}

func checkError(err error) {
	// Função para checar erros
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func CreateConn(serverHost string, serverPort int) *net.UDPConn {
	// Retorna um endereço de udp
	udpAddr, err := net.ResolveUDPAddr("udp", serverHost+":"+strconv.Itoa(serverPort))
	checkError(err)

	// Cria o socket UDP
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	return conn
}

func NewServerRequestHandlerUDP(serverHost string, serverPort int) ServerRequestHandlerUDP {
	// Cria um novo ServerRequestHandlerUDP
	srh := new(ServerRequestHandlerUDP)
	srh.serverHost = serverHost
	srh.serverPort = serverPort
	srh.conn = CreateConn(srh.serverHost, srh.serverPort)

	return *srh
}

func (srh *ServerRequestHandlerUDP) Receive() []byte {
	var err error
	var n int

	// Recebe a mensagem
	msgFromClient := make([]byte, 8192)
	n, srh.addr, err = srh.conn.ReadFromUDP(msgFromClient)
	checkError(err)

	return msgFromClient[:n]
}

func (srh *ServerRequestHandlerUDP) Send(msg []byte) {

	// Envia a mensagem
	_, err := srh.conn.WriteTo(msg, srh.addr)
	checkError(err)
}
