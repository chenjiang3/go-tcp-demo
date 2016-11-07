package main

import (
	"net"
	"log"
	"fmt"
	"common"
	"time"
)

func handle()  {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
	var count = 0
//	conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	for {
		count++
		fmt.Println("before send: count = ", count)
		_, err := conn.Write(common.NewHeader(int32(count), int32(count), 5, 2).ToBytes())
		if err != nil {
			fmt.Println("write error ", err)
			return
		}
		fmt.Println("after send: count = ", count)

//		time.Sleep(time.Second * 2)
	}
}

func main() {

	done := make(chan bool)

	go handle()

	for true {
		time.Sleep(time.Second)
	}

	//for {
	//	time.Sleep(time.Second)
	//}

	<- done

}
