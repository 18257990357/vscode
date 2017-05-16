package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func sender(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "./run.sh")
	words, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}
	conn.Write([]byte(words))
	fmt.Println("send over")
}

func main() {
	server := "211.140.139.66:1984"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

}
