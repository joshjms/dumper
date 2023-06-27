package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8001\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		url := "https://test.gtp-trial.teleport.sh/query/"

		// GraphQL query
		query := `
			query {
				getAllFruits{
					id,
					name
				}
			}
		`

		// Create an HTTP client
		client := &http.Client{}

		// Create a JSON payload with the query
		payload := map[string]interface{}{
			"query": query,
		}

		// Convert payload to JSON
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Failed to marshal JSON payload:", err)
			return
		}

		// Create an HTTP POST request with the payload
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
		if err != nil {
			fmt.Println("Failed to create request:", err)
			return
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Failed to send request:", err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}

		// Print the response
		fmt.Println("Response:", string(body))

		fmt.Fprintf(w, "Response: %v\n", string(body))
	})

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
