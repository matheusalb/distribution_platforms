package clientrequesthandlertcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

type ClientRequestHandlerTCP struct {
	HostToConn string
	PortToConn int
	clientConn net.Conn
}

func checkError(err error) {
	// Função para checar erros

	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func NewClientRequestHandlerTCP(hostToConn string, portToConn int) *ClientRequestHandlerTCP {
	// Função para criar o ClientRequestHandlerTCP
	crh := new(ClientRequestHandlerTCP)
	crh.HostToConn = hostToConn
	crh.PortToConn = portToConn
	crh.clientConn = nil

	return crh
}

func (crh *ClientRequestHandlerTCP) Connection() {
	// Função para criar a conexão com o servidor
	isActive := func(c net.Conn) bool {
		one := []byte{}
		c.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
		_, err := c.Read(one)
		if err == io.EOF {
			return false
		}
		c.SetReadDeadline(time.Time{}) // Remover o timeout
		return true
	}

	if crh.clientConn != nil {
		if isActive(crh.clientConn) {
			return
		}
	}

	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", crh.HostToConn+":"+strconv.Itoa(crh.PortToConn))
		if err == nil {
			break
		}
	}
	crh.clientConn = conn
}

func (crh *ClientRequestHandlerTCP) Send(msg []byte) {
	// Função para enviar mensagem para o servidor

	// Cria conexão com servidor
	crh.Connection()

	// Cria um slice de bytes com o tamanho da mensagem
	sizeMsgToServer := make([]byte, 4)
	len := uint32(len(msg))
	binary.LittleEndian.PutUint32(sizeMsgToServer, len)

	// Envia o tamanho da mensagem
	_, err := crh.clientConn.Write(sizeMsgToServer)
	checkError(err)

	// Envia a mensagem de fato
	_, err = crh.clientConn.Write(msg)
	checkError(err)
}

func (crh *ClientRequestHandlerTCP) Receive() []byte {
	// Recebe o tamanho da msg
	sizeMsgFromServer := make([]byte, 4)
	_, err := crh.clientConn.Read(sizeMsgFromServer)
	checkError(err)

	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	// Recebe a mensagem de fato
	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = crh.clientConn.Read(msgFromServer)
	checkError(err)

	return msgFromServer
}

func (crh *ClientRequestHandlerTCP) Close() {
	// Função para fechar a conexão com o servidor
	crh.clientConn.Close()
	crh.clientConn = nil
}
