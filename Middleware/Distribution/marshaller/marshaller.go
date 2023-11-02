package marshaller

import (
	"Middleware/Distribution/miop"
	"encoding/json"
	"log"
)

type Marshaller struct{}

func (Marshaller) Marshall(msg miop.Packet) []byte {
	r, err := json.Marshal(msg)

	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

func (Marshaller) Unmarshall(msg []byte) miop.Packet {
	packet := miop.Packet{}

	err := json.Unmarshal(msg, &packet)

	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}

	return packet
}
