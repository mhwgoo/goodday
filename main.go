package main

import (
	"fmt"
	"os"
	"time"
)

var user = os.Getenv("USER")
var tasks []string
var reads []string
var lessons []string
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

	tasks = []string{"SPEAK: sentences", "GO: godl and git push", "GO: achieve and git push", "KATE: git push"}
	print("TO DO", tasks)

	reads = []string{"The Apology", "The Guardian", "bash arith"}
	print("TO READ", reads)

	lessons = []string{"socket", "docker", "bash scripting", "w3m"}
	print("TO LEARN", lessons)

	problems = []string{"emerge log", "dual monitor support", "visualizer display"}
	print("TO SOLVE", problems)
}

func print(title string, items []string) {
	fmt.Printf("\n\n%s:", title)
	for index, item := range items {
		fmt.Printf("\n%d %s", index+1, item)
	}
}
