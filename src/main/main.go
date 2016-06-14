package main

import (
	"net/http"
	"io/ioutil"
)

func main() {
	hand := new(MyHandler)
	http.Handle("/", hand)
	http.ListenAndServe(":3000",nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}