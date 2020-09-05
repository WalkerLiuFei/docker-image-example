/*
  author Walker
  created at 2020/9/5 15:08
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func doHttpRequest() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			response, err := http.Get("http://http-server-service:8888/test")
			if err != nil || response == nil || response.Body == nil {
				fmt.Printf("error when do request %s\n", err.Error())
				continue
			}
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("read from response body failed : %s", err.Error())
				continue
			}
			fmt.Println(string(responseData))
			response.Body.Close()
		}
	}
}
