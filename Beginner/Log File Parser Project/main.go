package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// parseLogFile opens a log file, reads it line by line, and counts occurrences of specified keywords
func parseLogFile(filename string) (map[string]int, error) {
    // Initialize a map to store counts for each keyword
    counts := map[string]int{"INFO": 0, "ERROR": 0, "WARNING": 0}
    
    // Open the log file
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Use a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        
        // Check for each keyword in the line and increment the respective count
        for key := range counts {
            if strings.Contains(line, key) {
                counts[key]++
            }
        }
    }

    // Check for any errors encountered during scanning
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return counts, nil
}

func main() {
    // Specify the log file path (change this path to your actual log file)
    logFile := "example.log"

    // Parse the log file and get the counts for each keyword
    counts, err := parseLogFile(logFile)
    if err != nil {
        fmt.Println("Error reading log file:", err)
        return
    }

    // Display the results
    fmt.Printf("Log analysis for %s:\n", logFile)
    for key, count := range counts {
        fmt.Printf("%s: %d\n", key, count)
    }
}
