package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
)

const delimiter = '\n'

func accept(listener net.Listener) {
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Printf("Connect to %s \n", conn.RemoteAddr())
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for true {
		data, err := reader.ReadBytes(delimiter)
		if err != nil {
			fmt.Printf("Error read data ：%s", err.Error())
			conn.Close()
			return
		}
		fmt.Printf("from client : %s ", string(data))
		io.WriteString(conn, "Received\n")
	}
}
func main() {
	fmt.Printf("max CPU core number is %d\n", runtime.NumCPU())
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server start successful")
	accept(listener)
}
