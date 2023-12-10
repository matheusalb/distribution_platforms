package invoker

import (
	"Middleware/Distribution/lifecyclemanagement"
	"Middleware/Distribution/marshaller"
	"Middleware/Distribution/miop"
	"Middleware/Infrastructure/serverrequesthandlertcp"
	shared "Middleware/Shared"
	"fmt"
	"sync"
)

type Invoker struct {
	Port int
}

var lcm lifecyclemanagement.LifecycleManager
var mutex sync.Mutex

func handle_servant(servant *lifecyclemanagement.Servant, srh serverrequesthandlertcp.ServerRequestHandlerTCP) {

	marshall := marshaller.Marshaller{}
	params := make([]interface{}, 1)
	for {
		msgBytes := srh.Receive()

		pck := marshall.Unmarshall(msgBytes)
		op := pck.PackBody.Msg.HeaderMsg.Operation

		close := false
		switch op {
		case "DownloadBook":
			servant.Update_life()
			book := servant.Impl.DownloadBook(pck.PackBody.Msg.BodyMsg.Body[0].(string))
			params[0] = book
		case "Close":
			close = true
		}

		if !close {
			msgHeader := miop.MessageHeader{
				Context: "Response", RequestId: 0, Status: 1,
			}

			msgBody := miop.MessageBody{Body: params}

			packHeader := miop.PacketHeader{Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REPLY, Magic: "MIOP"}
			packBody := miop.PacketBody{Msg: miop.Message{HeaderMsg: msgHeader, BodyMsg: msgBody}}

			pckg := miop.Packet{PackHeader: packHeader, PackBody: packBody}

			msgToClientBytes := marshall.Marshall(pckg)

			srh.Send(msgToClientBytes)
		}

		if close || servant.IsExpired() {
			fmt.Println("Expirou")
			srh.CloseConn()
			break
		}
	}

	mutex.Lock()
	lcm.ReturnServant(servant)
	mutex.Unlock()
}

func (invoker Invoker) Invoke() {
	srh := serverrequesthandlertcp.NewServerRequestHandlerTCP(shared.N_HOST_SERVIDOR, invoker.Port)

	lcm := lifecyclemanagement.NewLifecycleManager(10, 5)
	lcm.Pooling()

	go lcm.Leasing()

	for {
		srh.Accept() // Aceita conexão...

		servant := lcm.GetServant()
		fmt.Println(servant)
		go handle_servant(servant, srh)
	}
}
