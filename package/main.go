package main

import (
	"log"
	"net/http"
	_ "package/router"
)

//import "fmt"

func main() {
	//	code, _ := upload("files/1.txt", "files/2.txt")
	//	fmt.Println(code) //list the file
	//path("./")
	//设置路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
