package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "path to package list")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: ./sorting --file <path_to_package_list>")
		os.Exit(1)
	}

	// Read file
	content, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Split lines, remove empty/whitespace-only lines
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %v\n", err)
		os.Exit(1)
	}

	// Remove duplicates
	unique := make(map[string]struct{})
	var result []string
	for _, line := range lines {
		if _, exists := unique[line]; !exists {
			unique[line] = struct{}{}
			result = append(result, line)
		}
	}

	// Sort
	sort.Strings(result)

	// Write back to file
	err = os.WriteFile(*filePath, []byte(strings.Join(result, "\n")+"\n"), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}
