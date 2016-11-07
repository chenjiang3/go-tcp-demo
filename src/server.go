package main

import (
	"net"
	"log"
	"io"
	"common"
	"fmt"
)

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Println("\n\n\n")
	count := 0
	for {
		var buf = make([]byte, common.HeaderLength)
		_, err := io.ReadFull(c, buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}

		header := common.ReadHeader(buf)
		count++
		fmt.Printf("receive packet count %d: ", count)
		header.Description()
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Println("listen error: ", err)
		return
	}

	count := 0

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			break
		}

		count++
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>accept a new connection, client count = \n", count)

		go handleConn(conn)
	}
}
