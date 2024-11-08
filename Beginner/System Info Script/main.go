package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    // Fetching OS and Architecture
    fmt.Println("Operating System:", runtime.GOOS)
    fmt.Println("Architecture:", runtime.GOARCH)

    // Fetching Hostname
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println("Error fetching hostname:", err)
    } else {
        fmt.Println("Hostname:", hostname)
    }

    // Fetching CPU count
    fmt.Println("Available CPUs:", runtime.NumCPU())

    // Fetching Go version
    fmt.Println("Go Version:", runtime.Version())
}
