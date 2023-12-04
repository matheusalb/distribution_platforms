package serverrequesthandlertcp

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
)

type ServerRequestHandlerTCP struct {
	serverHost     string
	serverPort     int
	ListenerServer net.Listener
	conn           net.Conn
}

func checkError(err error) {
	// Função para checar erros

	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func CreateListen(serverHost string, serverPort int) net.Listener {
	// Função para criar o Listener TCP
	ListenerServer, err := net.Listen("tcp", serverHost+":"+strconv.Itoa(serverPort))
	checkError(err)
	return ListenerServer
}

func NewServerRequestHandlerTCP(serverHost string, serverPort int) ServerRequestHandlerTCP {
	// Função para criar o ServerRequestHandlerTCP
	srh := new(ServerRequestHandlerTCP)
	srh.serverHost = serverHost
	srh.serverPort = serverPort
	srh.ListenerServer = CreateListen(srh.serverHost, srh.serverPort)
	srh.conn = nil

	return *srh
}

func (srh *ServerRequestHandlerTCP) Accept() net.Conn {
	// Função para aceitar conexão do cliente
	conn, err := srh.ListenerServer.Accept()
	checkError(err)

	return conn
}

func (srh *ServerRequestHandlerTCP) Receive() []byte {
	// Função para receber mensagem do cliente

	if srh.conn == nil {
		srh.conn = srh.Accept()
	}

	sizeMsgFromClient := make([]byte, 4)
	_, err := srh.conn.Read(sizeMsgFromClient)
	checkError(err)
	len := binary.LittleEndian.Uint32(sizeMsgFromClient)

	msgFromClient := make([]byte, len)
	_, err = srh.conn.Read(msgFromClient)
	checkError(err)

	return msgFromClient
}

func (srh *ServerRequestHandlerTCP) Send(msg []byte) {
	// Função para enviar mensagem para o cliente

	// Cria um slice de bytes com o tamanho da mensagem
	sizeMsgToClient := make([]byte, 4)
	len := uint32(len(msg))
	binary.LittleEndian.PutUint32(sizeMsgToClient, len)

	// Envia o tamanho da mensagem
	_, err := srh.conn.Write(sizeMsgToClient)
	checkError(err)

	// Envia a mensagem
	_, err = srh.conn.Write(msg)
	checkError(err)

}

func (srh *ServerRequestHandlerTCP) CloseConn() {
	srh.conn.Close()
	srh.conn = nil
}
