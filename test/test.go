package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	receiver, err := net.Listen("tcp", "127.0.0.1:9595")
	if err != nil {
		panic(err)
	}
	defer receiver.Close()

	for {
		cC, err := receiver.Accept()
		if err != nil {
			panic(err)
		}
		go server(cC)
	}
}

func server(c net.Conn) {
	var size uint32
	err := binary.Read(c, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	c.SetReadDeadline(time.Now().Add(5 * time.Second))

	msgSize := make([]byte, size)
	_, err = c.Read(msgSize)
	if err != nil {
		panic(err)
	}

	msgStrg := string(msgSize)
	fmt.Printf("Received Connection, Message: %s\n", msgStrg)
}
