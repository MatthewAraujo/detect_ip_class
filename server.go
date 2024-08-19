package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ipClass struct {
	class string
	min   int
	max   int
}

func address_ip(message string) string {
	classes := [4]ipClass{
		{"A", 0, 127},
		{"B", 128, 191},
		{"C", 192, 223},
		{"D", 224, 239},
	}

	fb, err := strconv.Atoi(strings.Split(message, ".")[0])
	if err != nil {
		return "Invalid IP address \n"
	}

	target := ipClass{}
	for _, c := range classes {
		if c.min <= fb && c.max >= fb {
			target.class = c.class
			break
		}
	}

	return fmt.Sprintf("class %s \n", target.class)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			return
		}

		message = address_ip(message)

		response := strings.ToUpper(message)
		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(conn)
	}
}
