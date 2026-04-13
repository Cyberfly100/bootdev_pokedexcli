package main

func cleanInput(text string) []string {
	result := make([]string, 0)
	currentWord := ""
	for _, char := range text {
		if char > 32 && char < 127 {
			if char >= 'A' && char <= 'Z' {
				char += 32 // Convert to lowercase
			}
			currentWord += string(char)
		} else {
			if currentWord != "" {
				result = append(result, currentWord)
				currentWord = ""
			}
		}
	}
	if currentWord != "" {
		result = append(result, currentWord)
	}
	return result
}
