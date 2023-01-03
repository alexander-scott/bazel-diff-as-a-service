// Package validation validates received requests
package validation

import (
	"encoding/json"
	"fmt"
)

// RequestData contains information about the request
type RequestData struct {
	GitURL      string `json:"git_url"`
	StartCommit string `json:"start_commit"`
	EndCommit   string `json:"end_commit"`
}

// ValidateRequest returns variables from request body
func ValidateRequest(bytes []byte) RequestData {
	// Unmarshal the JSON string into a dict
	var requestData RequestData
	error := json.Unmarshal(bytes, &requestData)
	if error != nil {
		fmt.Println(error)
	}

	// Check if the two necessary keys are present and print them
	fmt.Printf("Git URL: %s\n", requestData.GitURL)
	fmt.Printf("Starting hash: %s\n", requestData.StartCommit)
	fmt.Printf("Ending hash: %s\n", requestData.EndCommit)

	return requestData
}
