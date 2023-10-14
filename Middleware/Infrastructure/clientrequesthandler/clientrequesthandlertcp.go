package clientrequesthandlertcp

import (
	"encoding/binary"
	"net"
	"strconv"
	"fmt"
)

type ClientRequestHandlerTCP struct {
	name string
	hostToConn string
	portToConn int
	conn net.Conn
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func NewClientRequestHandlerTCP(name string, hostToConn string, portToConn int) *ClientRequestHandlerTCP {
	crh = new(ClientRequestHandlerTCP)
	crh.name = name
	crh.hostToConn = hostToConn
	crh.portToConn = portToConn
	crh.conn = nil

	return crh
}

func (crh *ClientRequestHandler) Connection() {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", crh.hostToConn+":"+strconv.Itoa(crh.portToConn))
		if err == nil {
			break
		}
	}
	crh.clientConn = conn
}


func (crh *ClientRequestHandlerTCP) Send(msg string) {
	if crh.conn == nil {
		crh.Connection()
	}

	sizeMsgToServer := make([]byte, 4)
	len := uint32(len(msg))
	binary.LittleEndian.PutUint32(sizeMsgToServer, len)

	// Envia o tamanho da mensagem
	_, err = crh.conn.Write(sizeMsgToServer)
	checkError(err)


	// Envia a mensagem de fato
	_, err = crh.conn.Write([]byte(msg))
	checkError(err)
}

func (crh *ClientRequestHandlerTCP) Receive() string {
	// Recebe o tamanho da msg
	sizeMsgFromServer := make([]byte, 4)
	_, err := crh.clientConn.Read(sizeMsgFromServer)
	checkError(err)

	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = crh.clientConn.Read(msgFromServer)
	checkError(err)

	return msgFromServer
}

