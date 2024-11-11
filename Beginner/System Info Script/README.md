# System Info Script

This Go script fetches and displays basic system information, such as the operating system, architecture, hostname, CPU count, and Go version.

## Features

- **Operating System**: Displays the operating system name.
- **Architecture**: Shows the system's architecture (e.g., `amd64`, `arm`).
- **Hostname**: Displays the hostname of the machine.
- **CPU Count**: Shows the number of available CPUs.
- **Go Version**: Displays the currently installed Go version.

## Requirements

- Go must be installed. You can download it from [golang.org](https://golang.org/dl/).

## Usage

1. **Clone or download** the project to your local machine.
2. **Navigate** to the directory containing the `system_info.go` file.
3. Run the following command in the terminal:

   ```bash
   go run system_info.go
   ```


## Example Output

 ```bash

    Operating System: linux
    Architecture: amd64
    Hostname: your-hostname
    Available CPUs: 4
    Go Version: go1.19.1
```
## Explanation of the Script
OS and Architecture: Uses runtime.GOOS and runtime.GOARCH to fetch the operating system and system architecture.
Hostname: Uses os.Hostname() to retrieve the hostname of the system.
CPU Count: Uses runtime.NumCPU() to get the number of available CPUs.
Go Version: Uses runtime.Version() to display the Go runtime version.

   
