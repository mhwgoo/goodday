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
	name   string
	status statusID
}

type statusID int

const (
	STATUS_TODO statusID = iota
	STATUS_DOING
	STATUS_DONE
)

func main() {
	codes := task{
		name:  "TO CODE",
		items: []item{{"go: godl", STATUS_DOING}, {"go: achieve", STATUS_DONE}, {"go: socket", STATUS_TODO}},
	}

	reads := task{
		name:  "TO READ",
		items: []item{{"speak: sentence", STATUS_TODO}, {"The Apology", STATUS_DOING}, {"The Guardian", STATUS_TODO}},
	}

	problems := task{
		name:  "TO SOLVE",
		items: []item{{"virtualbox-guest-addons", STATUS_TODO}, {"dual monitor", STATUS_TODO}},
	}

	print(codes.name, codes.items)
	print(reads.name, reads.items)
	print(problems.name, problems.items)
}

func print(title string, items []item) {
	fmt.Printf("\n%s:\n", title)
	for key, value := range items {
		switch value.status {
		case 0:
			fmt.Printf("%d %s [ ]\n", key+1, value.name)
		case 1:
			fmt.Printf("%d %s [=]\n", key+1, value.name)
		case 2:
			fmt.Printf("%d %s [x]\n", key+1, value.name)
		}
	}
}

func init() {
	dt := time.Now()
	fmt.Println(dt.Format("Mon Jan 02 15:04"))
	hour := dt.Hour()
	if hour >= 0 && hour <= 3 {
		fmt.Printf("Good Night, %s :)\n", user)
	} else if hour > 3 && hour < 12 {
		fmt.Printf("Good Morning, %s :)\n", user)
	} else if hour == 12 {
		fmt.Printf("Good Noon, %s :)\n", user)
	} else if hour > 12 && hour <= 15 {
		fmt.Printf("Good Afternoon, %s :)\n", user)
	} else if hour > 15 && hour < 20 {
		fmt.Printf("Good Evening, %s :)\n", user)
	} else if hour >= 20 && hour < 24 {
		fmt.Printf("Good Night, %s :)\n", user)
	}
}
