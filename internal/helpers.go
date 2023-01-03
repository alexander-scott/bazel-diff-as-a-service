// Package internal provides simple helper functions
package internal

import "strings"

// EscapeStringBeforeLogging removes bad characters from strings
func EscapeStringBeforeLogging(inputString string) string {
	escapedString := strings.ReplaceAll(inputString, "\n", "")
	escapedString = strings.ReplaceAll(escapedString, "\r", "")
	return escapedString
}
