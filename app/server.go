package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func readRequest(conn net.Conn) {

	// read the request and store it
	var buf []byte = make([]byte, 1024)

	_, err := conn.Read(buf) // wait till client sends request blocking call
	if err != nil {
		log.Fatal(err)
	}

	// emulate - do some process
	log.Println("processing the request")
	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	conn.Close()

}

func main() {
	fmt.Println("starting http server......")

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("error during listing on port")
		log.Fatal(err)
	}

	for {

		// Wait and listen for a connection
		log.Println("waiting for client to connect")
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error during waiting for a connection")
			log.Fatal(err)
		}

		log.Println("client connected")

		go readRequest(conn)

	}
}
