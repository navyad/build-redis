package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)



func handelConnection(conn net.Conn) {
	for {
		buf := make([]byte, 124)
		arrayLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		data := string(buf[:arrayLen])
		if strings.Contains(data, "ping") {
			conn.Write([]byte("+PONG\r\n"))
		}
        }
}


func main() {
	tcp_listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port: 6379")
		os.Exit(1)
	}
	fmt.Println("Accepting connection")

	conn, err := tcp_listener.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	} else {
		conn.Write([]byte("+PONG\r\n"))
	}
	handelConnection(conn)
}
