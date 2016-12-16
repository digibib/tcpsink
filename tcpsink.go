package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strings"
)

var (
	listenPort = flag.String("l", "9999", "Port to listen on")
	listenHost = flag.String("h", "localhost", "Host to listen on")
	prefix     = flag.String("p", "tcpsink: ", "String to prefix log output")
	verbose    = flag.Int("v", 0, "Verbosity level")
	strip      = flag.String("s", "", "remove line containing this string")
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
	if *strip != "" {
		log.Printf("Removing lines containing %s", *strip)
	}
	for {
		// incoming connections.
		c, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
		}
		if *verbose > 0 {
			log.Printf("[%s] --> CONNECTED", c.RemoteAddr().String())
		}
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
		if strings.Contains(s, *strip) {
			if *verbose > 0 {
				log.Printf("Removed line containing %s", *strip)
			}
		} else {
			log.Printf("[%s] %s", c.RemoteAddr().String(), s)
		}
	}
	if *verbose > 0 {
		log.Printf("[%s] <-- DISCONNECTED", c.RemoteAddr().String())
	}
	c.Close()
}
