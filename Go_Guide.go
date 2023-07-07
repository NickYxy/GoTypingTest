package GoTypingTest

import (
	"errors"
	"fmt"
	"io"
)

type Packet struct {
	PacketType    uint8
	PacketVersion uint8
	Data          *Data
}

type Data struct {
	Stat uint8
	Len  uint8
	Buf  [8]byte
}

func (p *Packet) UnmarshalBinary(b []byte) error {
	if len(b) < 2 {
		return io.EOF
	}

	p.PacketType = b[0]
	p.PacketVersion = b[1]

	if len(b) > 2 {
		p.Data = new(Data)
	}

	return nil
}

func foo(c chan int) {
	defer close(c)
	_, err := parse(1, []byte("haha"))

	if err != nil {
		c <- 0
		return
	}
	c <- 1
}

func parse(lenControlByUser int, data []byte) ([]byte, error) {
	size := lenControlByUser
	if size > 64*1024*1024 {
		return nil, errors.New("value too large")
	}

	buffer := make([]byte, size)
	copy(buffer, data)
	return buffer, nil
}

func main() {
	packet := new(Packet)
	data := make([]byte, 2)

	if err := packet.UnmarshalBinary(data); err != nil {
		fmt.Println("Failed to unmarshall packet")
		return
	}

	if packet.Data == nil {
		return
	}

	fmt.Printf("Stat: %v\n", packet.Data.Stat)
}
