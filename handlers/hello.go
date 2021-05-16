package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) (h *Hello) {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", data)

}
