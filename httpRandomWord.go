package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func randomWord() string {
	resp, err := http.Get("https://random-word-api.vercel.app/api?words=1")
	if err != nil {
		log.Printf("error sending request: %v", err )
		return "Error getting word"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading answer: %v", err)
		return "Error reading word"
	}
	
	
	var result []string
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Printf("Error unmarshalling json: %v", err)
        return "Error parsing word"
    }

	return result[0]
}	