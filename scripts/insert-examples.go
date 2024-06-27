package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	templateContent, err := os.ReadFile("README.template")
	if err != nil {
		fmt.Println("Error reading README.template file:", err)
		panic(err)
	}

	// Replace %insert path/to/file.go% with the content of the file
	re := regexp.MustCompile(`%insert ([^%]+)%`)

	// Replace the placeholders with the file contents
	result := re.ReplaceAllFunc(templateContent, func(match []byte) []byte {
		// Extract the file path from the placeholder
		matches := re.FindSubmatch(match)
		if len(matches) < 2 {
			return match
		}
		filePath := string(matches[1])

		// Read the content of the file to be inserted
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			panic(err)
		}

		return fileContent
	})

	// Write the result to a new file or overwrite the existing file
	err = os.WriteFile("README.md", result, 0644)
	if err != nil {
		fmt.Println("Error writing README.md file:", err)
		panic(err)
	}

	fmt.Println("README.md file created successfully.")
}
