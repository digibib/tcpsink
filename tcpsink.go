package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var (
	listenPort = flag.String("l", "9999", "Port to listen on")
	listenHost = flag.String("h", "localhost", "Host to listen on")
	prefix     = flag.String("p", "tcpsink: ", "String to prefix log output")
)

func main() {
	flag.Parse()
	log.SetPrefix(*prefix)

	l, err := net.Listen("tcp", *listenHost+":"+*listenPort)
	if err != nil {
		log.Fatalf("Error listening:", err.Error())
	}

	defer l.Close()

	log.Printf("Listening on %s:%s", *listenHost, *listenPort)
	for {
		// incoming connections.
		c, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
		}
		log.Printf("[%s] --> CONNECTED", c.RemoteAddr().String())
		go handleRequest(c)
	}
}

func handleRequest(c net.Conn) {

	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading: %s", err.Error())
			}
			break
		}
		s := string(buf[:n])
		log.Printf("[%s] %s", c.RemoteAddr().String(), s)
	}
	log.Printf("[%s] <-- DISCONNECTED", c.RemoteAddr().String())
	c.Close()
}
