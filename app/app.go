package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Send message")
		fmt.Println("2. Exit")
		scanner.Scan()
		option := scanner.Text()
		if option == "1" {
			sendMessage()
		} else if option == "2" {
			fmt.Println("Thanks, Goodbye!")
			break
		}
	}
}

func sendMessage() {
	scanner := bufio.NewScanner(os.Stdin)
	var message string
	scanner.Scan()
	message = scanner.Text()

	sC, err := net.DialTimeout("tcp", "127.0.0.1:9595", 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer sC.Close()

	sC.SetWriteDeadline(time.Now().Add(5 * time.Second))

	err = binary.Write(sC, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}

	_, err = sC.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func main() {
	menu()
}
