package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
)

func accept(listener net.Listener) {
	defer listener.Close()
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
		data, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error read data ：%s", err.Error())
			return
		}
		fmt.Printf("from client : %s ", string(data))
		n, err := io.WriteString(conn, "Received\n")
		if err != nil {
			fmt.Printf("Error write data ：%s,size : %d", err.Error(), n)
			return
		}
	}
}
func main() {
	certificate, err := tls.LoadX509KeyPair("./asset/server.crt", "./asset/server.key")
	if err != nil {
		fmt.Printf("Load certificate error %s \n", err.Error())
	}
	config := &tls.Config{Certificates: []tls.Certificate{certificate}}
	listener, err := tls.Listen("tcp", ":8888", config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Start server success")
	accept(listener)
}
