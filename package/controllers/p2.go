package controllers

import (
	"fmt"
	"net/http"
)
func init(){
	
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
func test() {
	fmt.Println("test")
}
