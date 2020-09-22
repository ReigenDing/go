package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	if connection, e := net.Dial("tcp", ":1024"); e == nil {
		defer connection.Close()
		if text, e := bufio.NewReader(connection).ReadString('\n'); e == nil {
			fmt.Println(text)
		}
	}
}
