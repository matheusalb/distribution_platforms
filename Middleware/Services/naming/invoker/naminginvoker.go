package invoker

import (
	"Middleware/Distribution/clientproxy"
	"Middleware/Distribution/marshaller"
	"Middleware/Distribution/miop"
	"Middleware/Infrastructure/serverrequesthandlertcp"
	"Middleware/Services/naming"
	shared "Middleware/Shared"
)

type NamingInvoker struct{}

func (NamingInvoker) Invoke() {
	srhImpl := serverrequesthandlertcp.NewServerRequestHandlerTCP("localhost", shared.N_PORT_NS)
	marshallerImpl := marshaller.Marshaller{}
	namingImpl := naming.NamingService{}
	miopPacketReply := miop.Packet{}
	replyParams := make([]interface{}, 1)

	// control loop
	for {
		// receive request packet
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall request packet
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)

		// extract operation name
		operation := miopPacketRequest.PackBody.Msg.HeaderMsg.Operation

		// demux request
		switch operation {
		case "Register":
			_p1 := miopPacketRequest.PackBody.Msg.BodyMsg.Body[0].(string)
			_map := miopPacketRequest.PackBody.Msg.BodyMsg.Body[1].(map[string]interface{})

			_proxyTemp := _map["Proxy"].(map[string]interface{})
			_p2 := clientproxy.ClientProxy{TypeName: _proxyTemp["TypeName"].(string), Host: _proxyTemp["Host"].(string), Port: int(_proxyTemp["Port"].(float64)), Id: int(_proxyTemp["Id"].(float64))}

			// dispatch request
			replyParams[0] = namingImpl.Register(_p1, _p2)
		case "Lookup":
			_p1 := miopPacketRequest.PackBody.Msg.BodyMsg.Body[0].(string)

			// dispatch request
			replyParams[0] = namingImpl.Lookup(_p1)
		}

		// assembly reply packet
		repHeader := miop.MessageHeader{Context: "", RequestId: miopPacketRequest.PackBody.Msg.HeaderMsg.RequestId, Status: 1}
		repBody := miop.MessageBody{Body: replyParams}
		header := miop.PacketHeader{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST}
		msg := miop.Message{HeaderMsg: repHeader, BodyMsg: repBody}
		body := miop.PacketBody{Msg: msg}
		miopPacketReply = miop.Packet{PackHeader: header, PackBody: body}

		// marshall reply packet
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send reply packet
		srhImpl.Send(msgToClientBytes)
		srhImpl.CloseConn()
	}
}
