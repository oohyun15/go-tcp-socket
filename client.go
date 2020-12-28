package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// client try to connect server
	conn, err := net.Dial("tcp", ":8000")

	if nil != err {
		log.Println(err)
	}

	// create goroutine
	// wait until server send a message then print it
	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	// block until client input then send a message to server
	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
