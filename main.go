package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Bad request submitted"))
			return
		}

		log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.ListenAndServe(":8080", nil)
}