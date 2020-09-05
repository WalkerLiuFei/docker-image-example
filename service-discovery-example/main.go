package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const delimiter = "\n"

func main() {
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
	waitChannel := make(chan int, 0)
	for _, address := range addresses {
		for count := 10; count >= 0; count++ {
			go func(address string) {
				go doConnect(address, port)
			}(address)
		}
	}
	<-waitChannel
}

func doConnect(address string, port int) {
	address = fmt.Sprintf("%s:%d", address, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", conn.RemoteAddr())
	doTicker(conn)
}

func doTicker(conn net.Conn) {
	ticker := time.NewTicker(time.Second)
	index := 0
	for range ticker.C {
		msg := fmt.Sprintf("this is a message %d %s", index, delimiter)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("write message failed %s\n", err.Error())
		}
		index++
	}
}
