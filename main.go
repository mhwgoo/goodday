package main

import (
	"fmt"
	"os"
	"time"
)

var user = os.Getenv("USER")

type task struct {
	name  kind
	items []item
}

type item struct {
	content string
	status  statusID
}

type statusID int

const (
	STATUS_TODO statusID = iota
	STATUS_DOING
	STATUS_DONE
)

type kind string

var taskKinds = []kind{"CODE", "READ", "SOLVE"}

func main() {
	codes := task{
		name:  taskKinds[0],
		items: []item{{"go: godl", STATUS_DOING}, {"go: achieve and learn from asana design concepts", STATUS_DONE}, {"go: socket", STATUS_TODO}},
	}

	reads := task{
		name:  taskKinds[2],
		items: []item{{"speak: sentence", STATUS_TODO}, {"The Apology", STATUS_DOING}, {"The Guardian", STATUS_TODO}},
	}

	problems := task{
		name:  taskKinds[2],
		items: []item{{"virtualbox-guest-addons", STATUS_TODO}, {"dual monitor", STATUS_TODO}},
	}

	print(codes.name, codes.items)
	print(reads.name, reads.items)
	print(problems.name, problems.items)
}

func print(name kind, items []item) {
	fmt.Printf("\n%s:\n", name)
	for key, value := range items {
		switch value.status {
		case 0:
			fmt.Printf("%d %s [ ]\n", key+1, value.content)
		case 1:
			fmt.Printf("%d %s [=]\n", key+1, value.content)
		case 2:
			fmt.Printf("%d %s [x]\n", key+1, value.content)
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
