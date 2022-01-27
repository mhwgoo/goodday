package main

import (
    "fmt"
    "time"
    "os"
)

var user = os.Getenv("USER")
var tasks []string

func main() {
    dt := time.Now()
    fmt.Println(dt.Format("Mon Jan 02 15:04"))
    hour := dt.Hour()
    if (hour > 0 && hour <= 3) {
        fmt.Printf("Good Night, %s :)", user)
    } else if (hour > 3 && hour < 12) {
        fmt.Printf("Good Morning, %s :)", user)
    } else if (hour == 12) {
        fmt.Printf("Good Noon, %s :)", user)
    } else if (hour > 12 && hour <= 15) {
        fmt.Printf("Good Afternoon, %s :)", user)
    } else if (hour > 15 && hour < 20) {
        fmt.Printf("Good Evening, %s :)", user)
    } else if (hour >= 20 && hour < 24) {
        fmt.Printf("Good Night, %s :)", user)
    }

    fmt.Println("\nToday's tasks:")
    tasks = []string {"read 'Go In Action'", "work on godl", "learn from Asana and build achieve", "read 'The Apology'"}
    for index, task := range tasks {
        fmt.Printf("%d %s\n", index + 1, task)
    }
}
