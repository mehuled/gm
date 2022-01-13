package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", gm)
	http.HandleFunc("/headers", headers)
	log.Println(http.ListenAndServe("0.0.0.0:8000", nil))
}

func gm(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "gm")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for key, val := range req.Header {
		fmt.Fprintf(w, "%s : %s \n", key, val)
	}
}
