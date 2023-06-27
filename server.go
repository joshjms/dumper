package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 9999\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	})

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
