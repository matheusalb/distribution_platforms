package requestor

import (
	aux "Middleware/Auxiliar"
	"Middleware/Distribution/marshaller"
	"Middleware/Distribution/miop"
	clientrequesthandlertcp "Middleware/Infrastructure/clientrequesthandlertcp"
	shared "Middleware/Shared"
)

// Requestor ...

type Requestor struct {
	crh *clientrequesthandlertcp.ClientRequestHandlerTCP
}

func (requestor *Requestor) initClientRequestHandlerTCP(Host string, Port int) {
	requestor.crh = clientrequesthandlertcp.NewClientRequestHandlerTCP(Host, Port)
}

func NewRequestor() Requestor {
	return Requestor{crh: nil}
}

// Invoke ...
func (requestor *Requestor) Invoke(inv aux.Invocation) []interface{} {

	marshall := marshaller.Marshaller{}

	if requestor.crh == nil {
		requestor.initClientRequestHandlerTCP(inv.Host, inv.Port)
	}

	msgHeader := miop.MessageHeader{
		Context: "Invoke", RequestId: 0,
		ResponseExpected: true, ObjectKey: 199,
		Operation: inv.Request.Op}

	msgBody := miop.MessageBody{Body: inv.Request.Params}

	packHeader := miop.PacketHeader{Version: "1.0", MessageType: shared.MIOP_REQUEST, Magic: "MIOP"}
	packBody := miop.PacketBody{Msg: miop.Message{HeaderMsg: msgHeader, BodyMsg: msgBody}}

	pckg := miop.Packet{PackHeader: packHeader, PackBody: packBody}

	marshalledPackage := marshall.Marshall(pckg)

	requestor.crh.Send(marshalledPackage)
	msg := requestor.crh.Receive()
	unmarshalledPackage := marshall.Unmarshall(msg)

	r := unmarshalledPackage.PackBody.Msg.BodyMsg.Body

	return r
}

func (request *Requestor) Close(inv aux.Invocation) {
	request.Invoke(inv)

	request.crh.Close()
	request.crh = nil
}
