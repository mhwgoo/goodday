package main

import (
	"fmt"
	"os"
	"time"
)

var user = os.Getenv("USER")
var tasks []string
var problems []string

func main() {
	dt := time.Now()
	fmt.Println(dt.Format("Mon Jan 02 15:04"))
	hour := dt.Hour()
	if hour >= 0 && hour <= 3 {
		fmt.Printf("Good Night, %s :)", user)
	} else if hour > 3 && hour < 12 {
		fmt.Printf("Good Morning, %s :)", user)
	} else if hour == 12 {
		fmt.Printf("Good Noon, %s :)", user)
	} else if hour > 12 && hour <= 15 {
		fmt.Printf("Good Afternoon, %s :)", user)
	} else if hour > 15 && hour < 20 {
		fmt.Printf("Good Evening, %s :)", user)
	} else if hour >= 20 && hour < 24 {
		fmt.Printf("Good Night, %s :)", user)
	}

	fmt.Println("\n\nToday's tasks:")

	tasks = []string{"learn socket", "work on godl", "work on achieve", "git commit kate repos", "read 'The Apology'"}
	for index, task := range tasks {
		fmt.Printf("%d %s\n", index+1, task)
	}

	fmt.Println("\nProblems to be solved:")
	problems = []string{"emerge log", "dual monitor support", "visualizer display"}
	for index, problem := range problems {
		fmt.Printf("%d %s\n", index+1, problem)
	}
}
