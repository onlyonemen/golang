// xss project main.go
package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/xss", xss)
	http.HandleFunc("/", code)
	http.ListenAndServe(":8080", nil)
}
func xss(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	//	w.Write([]byte(code))
	t, _ := template.ParseFiles("xss.html")

	t.Execute(w, code)
}
func code(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("code.html")

	t.Execute(w, nil)
}
