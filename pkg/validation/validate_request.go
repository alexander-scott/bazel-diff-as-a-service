// Package validation validates received requests
package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alexander-scott/bazel-diff-as-a-service/internal"
)

// RequestData contains information about the request
type RequestData struct {
	GitURL      string `json:"git_url"`
	StartCommit string `json:"start_commit"`
	EndCommit   string `json:"end_commit"`
}

// ValidateRequest returns variables from request body
func ValidateRequest(bytes []byte) (RequestData, error) {
	// Unmarshal the JSON string into a dict
	var requestData RequestData
	error := json.Unmarshal(bytes, &requestData)
	if error != nil {
		return RequestData{}, error
	}

	// Check if options are null
	if requestData.GitURL == "" || requestData.StartCommit == "" || requestData.EndCommit == "" {
		return RequestData{}, errors.New("invalid arguments provided")
	}

	// Check if the two necessary keys are present and print them
	fmt.Printf("Git URL: %s\n", internal.EscapeStringBeforeLogging(requestData.GitURL))
	fmt.Printf("Starting hash: %s\n", internal.EscapeStringBeforeLogging(requestData.StartCommit))
	fmt.Printf("Ending hash: %s\n", internal.EscapeStringBeforeLogging(requestData.EndCommit))

	return requestData, nil
}
