package main

import (
	"fmt"
	"net"
	"net/http"
)

const delimiter = '\n'

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		responseStr := fmt.Sprintf("request from %s,response from %s", request.RemoteAddr, GetLocalIP())
		writer.Write([]byte(responseStr))
	})
	http.ListenAndServe("0.0.0.0:8888", nil)
	fmt.Println("Server start successful")
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
