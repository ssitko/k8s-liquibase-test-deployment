package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Healthy response for 1.0.2!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
