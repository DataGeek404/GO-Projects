package main

import (
    "bufio"
    "cli-task-manager/task"
    "fmt"
    "os"
    "strconv"
)

func main() {
    tm := &task.TaskManager{}
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nTask Manager")
        fmt.Println("1. Add Task")
        fmt.Println("2. View Tasks")
        fmt.Println("3. Mark Task as Complete")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Exit")
        fmt.Print("Enter your choice: ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            fmt.Print("Enter task description: ")
            scanner.Scan()
            description := scanner.Text()
            tm.AddTask(description)
        case "2":
            tm.ViewTasks()
        case "3":
            fmt.Print("Enter task ID to mark complete: ")
            scanner.Scan()
            id, _ := strconv.Atoi(scanner.Text())
            tm.MarkComplete(id)
        case "4":
            fmt.Print("Enter task ID to delete: ")
            scanner.Scan()
            id, _ := strconv.Atoi(scanner.Text())
            tm.DeleteTask(id)
        case "5":
            fmt.Println("Exiting Task Manager. Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Try again.")
        }
    }
}
