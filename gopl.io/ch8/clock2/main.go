package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var flagPort = flag.Int("port", 9999, "put port for the server")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *flagPort))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening at localhost:%d\n", *flagPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
