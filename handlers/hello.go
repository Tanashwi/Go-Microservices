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

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//These 2 lines of code can be replaced with the one liner as below
		//rw.WriteHeader(http.StatusBadRequest)
		//rw.Write([]byte("Bad request submitted"))
		http.Error(rw, "Bad request submitted", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
