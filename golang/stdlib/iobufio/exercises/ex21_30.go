package exercises

import (
	"io"
	"log"
	"net"
	"os"
)

func Ex21_TCP_echo_server() {
	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("TCP server is listening at :8003...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			tc := io.TeeReader(c, os.Stdout)
			io.Copy(c, tc)
			c.Close()
		}(conn)
	}
}

// func Ex22_TCP_client() {

// }
