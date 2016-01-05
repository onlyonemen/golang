// Server project main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server start!")
	http.HandleFunc("/test", Test)
	http.ListenAndServe(":8001", nil)
}
func Test(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("ip")
	
	w.Write([]byte(ip))
}
