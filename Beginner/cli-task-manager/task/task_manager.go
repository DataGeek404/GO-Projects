package task

import (
    "fmt"
    "time" // Importing the time package
)

type TaskManager struct {
    Tasks []Task
}

func (tm *TaskManager) AddTask(description string) {
    task := Task{
        ID:          len(tm.Tasks) + 1,
        Description: description,
        Completed:   false,
        CreatedAt:   time.Now(),
    }
    tm.Tasks = append(tm.Tasks, task)
    fmt.Println("Task added:", description)
}

func (tm *TaskManager) ViewTasks() {
    for _, task := range tm.Tasks {
        status := "Pending"
        if task.Completed {
            status = "Completed"
        }
        fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, status)
    }
}

func (tm *TaskManager) MarkComplete(id int) {
    for i := range tm.Tasks {
        if tm.Tasks[i].ID == id {
            tm.Tasks[i].Completed = true
            fmt.Println("Task marked as complete:", tm.Tasks[i].Description)
            return
        }
    }
    fmt.Println("Task not found!")
}

func (tm *TaskManager) DeleteTask(id int) {
    for i, task := range tm.Tasks {
        if task.ID == id {
            tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
            fmt.Println("Task deleted:", task.Description)
            return
        }
    }
    fmt.Println("Task not found!")
}
