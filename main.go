package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello World!"

func main() {
	fmt.Println("starting main")
	
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
