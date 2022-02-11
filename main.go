package main

import (
	"fmt"
	"os"
	"time"
)

var user = os.Getenv("USER")

type task struct {
	name  string
	items []item
}

type item struct {
	name    string
	checked bool
}

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

	todos := task{
		name:  "TO DO",
		items: []item{{"speak: sentence", false}, {"go: godl", true}, {"go: achieve", false}},
	}

	reads := task{
		name:  "TO READ",
		items: []item{{"The Apology", false}, {"The Guardian", false}},
	}

	problems := task{
		name:  "TO SOLVE",
		items: []item{{"virtualbox-guest-addons", false}, {"dual monitor", false}},
	}

	print(todos.name, todos.items)
	print(reads.name, reads.items)
	print(problems.name, problems.items)
}

func print(title string, items []item) {
	fmt.Printf("\n\n%s:\n", title)
	for key, value := range items {
		if value.checked {
			fmt.Printf("%d %s [x]\n", key+1, value.name)
		} else {
			fmt.Printf("%d %s [ ]\n", key+1, value.name)
		}
	}
}
