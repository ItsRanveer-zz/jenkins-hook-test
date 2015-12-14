package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/testapi", testapi)

	log.Println("Listening at 3456...")

	err := http.ListenAndServe(":3456", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func testapi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Everyone!\n")
}
