package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// open socket
	l, err := net.Listen("tcp", ":8000")

	if nil != err {
		log.Println(err)
	}

	// close socket when main proc is done.
	defer l.Close()

	// connect socket and goroutine ConnHandler
	for {
		conn, err := l.Accept()
		if nil != err {
			log.Println(err)
			continue
		}
		defer conn.Close()
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		// blocked until client send a message
		// then read a message
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}

		// send client a message which client sent
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err = conn.Write(data[:n])
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
