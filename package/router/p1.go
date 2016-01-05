package router

import (
	//	"fmt"
	"net/http"
	"package/controllers"
)

func init() {
http.HandleFunc("/login", controllers.Hello) //设置访问的路由
}

//func login(w http.ResponseWriter, r *http.Request) {
//	//	fmt.Println("login")
//	controllers.hello(w, r)
//	controllers.test();
//}
