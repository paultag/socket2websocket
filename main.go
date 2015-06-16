package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:30010")
	if err != nil {
		fmt.Printf("Ouch, can't connect: %s", err)
		return
	}

	fmt.Fprintf(conn, "\n\n")

	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Ouch, error: %s", err)
			break
		}
	}
}
