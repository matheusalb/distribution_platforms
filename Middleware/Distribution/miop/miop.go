package miop

type MessageHeader struct {
	Context          string
	RequestId        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type MessageBody struct {
	Body []interface{}
}

type Message struct {
	HeaderMsg MessageHeader
	BodyMsg   MessageBody
}

// PacketHeader ...
type PacketHeader struct {
	Magic       string
	Version     string
	ByteOrder   bool
	MessageType int
	Size        string
}

type PacketBody struct {
	Msg Message
}

// Mensagem ficará dentro de um packet
type Packet struct {
	PackHeader PacketHeader
	PackBody   PacketBody
}
