/*
  author Walker
  created at 2020/9/5 14:09
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

func socketConnect() {
	// 直接DNS 拿对应服务地址
	serviceName := "echo-server-service.default.svc.cluster.local"
	addresses, err := net.LookupHost(serviceName)
	if err != nil {
		fmt.Printf("lookup address failed %s\n", err.Error())
	}
	/*port, err := net.LookupPort("tcp", serviceName)

	if err != nil {
		fmt.Printf("lookup port failed %s\n", err.Error())
	}*/
	port := 8888
	addressesBytes, err := json.Marshal(addresses)
	fmt.Println(string(addressesBytes))
	//waitChannel := make(chan int, 0)
	for _, address := range addresses {
		for count := 1; count >= 0; count++ {
			go func(address string) {
				doConnect(address, port)
			}(address)
		}
	}
	//<-waitChannel
}
func doConnect(address string, port int) {
	address = fmt.Sprintf("%s:%d", address, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("connect to echo server success : %s\n", conn.RemoteAddr())
	doTicker(conn)

}

func doTicker(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	go func() {
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
	}()
	ticker := time.NewTicker(time.Second)
	index := 0
	for range ticker.C {
		msg := fmt.Sprintf("this is a message %d \n", index)
		n, err := writer.Write([]byte(msg))
		fmt.Printf("msg [%s] write done %d\n", msg, n)
		if err != nil {
			fmt.Printf("write message failed %s\n", err.Error())
			conn.Close()
			return
		}
		index++
	}
}
