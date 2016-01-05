package web

import (

	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}
