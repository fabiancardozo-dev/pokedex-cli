package main

import "strings"

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	result := strings.Fields(lowerCaseText) // Split the string around white spaces
	return result
}
