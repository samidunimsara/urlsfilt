package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("i", "", "Input file containing URLs (required)")
	filters := flag.String("f", "", "Comma-separated filter patterns to exclude URLs")
	outputFile := flag.String("o", "", "Output file (optional, defaults to stdout)")
	help := flag.Bool("h", false, "Show help message")

	flag.Parse()

	// Show help if requested or no input file provided
	if *help || *inputFile == "" {
		showHelp()
		return
	}

	// Parse filter patterns
	var filterPatterns []string
	if *filters != "" {
		filterPatterns = strings.Split(*filters, ",")
		// Trim whitespace from each pattern
		for i := range filterPatterns {
			filterPatterns[i] = strings.TrimSpace(filterPatterns[i])
		}
	}

	// Read URLs from input file
	urls, err := readURLs(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Filter URLs
	filteredURLs := filterURLs(urls, filterPatterns)

	// Write output
	if err := writeOutput(filteredURLs, *outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	// Print summary to stderr
	fmt.Fprintf(os.Stderr, "Filtered %d URLs, %d remaining\n", len(urls)-len(filteredURLs), len(filteredURLs))
}

func showHelp() {
	fmt.Println("urlsfilt - URL Filtering Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  urlsfilt -i <input_file> [-f <filters>] [-o <output_file>]")
	fmt.Println("\nOptions:")
	fmt.Println("  -i string    Input file containing URLs (required)")
	fmt.Println("  -f string    Comma-separated filter patterns to exclude URLs")
	fmt.Println("  -o string    Output file (optional, defaults to stdout)")
	fmt.Println("  -h           Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  urlsfilt -i urls.txt -f www.")
	fmt.Println("  urlsfilt -i urls.txt -f www.,.js,outfit")
	fmt.Println("  urlsfilt -i urls.txt -f www. -o filtered.txt")
	fmt.Println("\nDescription:")
	fmt.Println("  Filter out URLs that contain any of the specified patterns.")
	fmt.Println("  Patterns are matched anywhere in the URL (domain, path, etc.).")
}

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func filterURLs(urls []string, patterns []string) []string {
	if len(patterns) == 0 {
		return urls
	}

	var filtered []string
	for _, url := range urls {
		if !shouldFilter(url, patterns) {
			filtered = append(filtered, url)
		}
	}

	return filtered
}

func shouldFilter(url string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(url, pattern) {
			return true
		}
	}
	return false
}

func writeOutput(urls []string, outputFile string) error {
	var writer *bufio.Writer
	var file *os.File

	if outputFile != "" {
		var err error
		file, err = os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()
		writer = bufio.NewWriter(file)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	for _, url := range urls {
		if _, err := writer.WriteString(url + "\n"); err != nil {
			return err
		}
	}

	return writer.Flush()
}
