package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	} else {
		fmt.Println("The address is ", addr.String())
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", addr.String(), 8080))
	if err != nil {
		fmt.Println("Invalid tcpAddr")
		os.Exit(1)
	}
	fmt.Println("tcpAddr ", tcpAddr.String())
	os.Exit(0)
}
