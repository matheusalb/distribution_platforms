package invoker

import (
	"Middleware/Distribution/marshaller"
	"Middleware/Distribution/miop"
	"Middleware/Infrastructure/serverrequesthandlertcp"
	shared "Middleware/Shared"
	"fmt"
	"io/ioutil"
	"os"
)

type Invoker struct{}

func readBook(nomeLivro string) string {
	file, err := os.Open("./books/" + nomeLivro + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}

func (Invoker) Invoke() {
	srh := serverrequesthandlertcp.NewServerRequestHandlerTCP(shared.N_HOST_SERVIDOR, shared.N_PORT_SERVIDOR)
	marshall := marshaller.Marshaller{}

	params := make([]interface{}, 1)
	for {
		msgBytes := srh.Receive()

		pck := marshall.Unmarshall(msgBytes)
		op := pck.PackBody.Msg.HeaderMsg.Operation

		switch op {
		case "Download":
			book := readBook(pck.PackBody.Msg.BodyMsg.Body[0].(string))
			params[0] = book
		}

		msgHeader := miop.MessageHeader{
			Context: "Response", RequestId: 0, Status: 1,
		}

		msgBody := miop.MessageBody{Body: params}

		packHeader := miop.PacketHeader{Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST, Magic: "MIOP"}
		packBody := miop.PacketBody{Msg: miop.Message{HeaderMsg: msgHeader, BodyMsg: msgBody}}

		pckg := miop.Packet{PackHeader: packHeader, PackBody: packBody}

		msgToClientBytes := marshall.Marshall(pckg)

		srh.Send(msgToClientBytes)
	}
}
