package requestor

import (
	"Middleware/Distribution/marshaller"
	"Middleware/Distribution/miop"
	clientrequesthandlertcp "Middleware/Infrastructure/clientrequesthandlertcp"
	shared "Middleware/Shared"
)

// Requestor ...

type Requestor struct{}

// Invoke ...
func (Requestor) Invoke(inv shared.Invocation) []interface{} {

	marshall := marshaller.Marshaller{}
	crh := clientrequesthandlertcp.NewClientRequestHandlerTCP(inv.Host, inv.Port)

	msgHeader := miop.MessageHeader{
		Context: "Invoke", RequestId: 0,
		ResponseExpected: true, ObjectKey: 199,
		Operation: inv.Request.Op}

	msgBody := miop.MessageBody{Body: inv.Request.Params}

	packHeader := miop.PacketHeader{Version: "1.0", MessageType: shared.MIOP_REQUEST, Magic: "MIOP"}
	packBody := miop.PacketBody{Msg: miop.Message{HeaderMsg: msgHeader, BodyMsg: msgBody}}

	pckg := miop.Packet{PackHeader: packHeader, PackBody: packBody}

	marshalledPackage := marshall.Marshall(pckg)

	crh.Send(marshalledPackage)
	msg := crh.Receive()
	unmarshalledMsg := marshall.Unmarshall(msg)

	r := unmarshalledMsg.PackBody.Msg.BodyMsg.Body

	return r
}
