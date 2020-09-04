package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	// 直接DNS 拿对应服务地址
	addresses, err := net.LookupAddr("echo-server-service.default.svc.cluster.local")
	if err != nil {
		fmt.Println(err.Error())
	}
	addressesBytes, err := json.Marshal(addresses)
	fmt.Println(string(addressesBytes))
}
