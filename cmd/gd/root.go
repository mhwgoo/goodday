package main

import (
	"database/sql"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mhwgoo/goodday"
	"github.com/spf13/cobra"
)

var user = os.Getenv("USER")

var DB, _ = sql.Open("sqlite3", "/home/kate/go/src/github.com/mhwgoo/goodday/goodday.db")
var Pool = goodday.NewPool(DB)

var rootCmd = &cobra.Command{
	Use: os.Args[0],
	Run: greet,
}

func greet(cmd *cobra.Command, args []string) {
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

	PrintQuotes()
	PrintSections(Sections)
}

func PrintSections(secs []Section) {
	for _, sec := range secs {
		fmt.Printf("\n[%s] %s\n", string(string(sec.Kind)[0]), strings.ToUpper(sec.Name))
		for _, task := range SortTasks(sec.Tasks) {
			if task.Priority == P1 {
				task.Priority = Level("***")
			} else if task.Priority == P2 {
				task.Priority = Level("*")
			} else {
				task.Priority = Level("")
			}
			switch task.Status {
			case Todo:
				fmt.Printf("  - [%s] %s %v [ ] %s\n", string(string(task.Kind)[0]), task.Name, task.Priority, task.Desc)
			case Doing:
				fmt.Printf("  - [%s] %s %v [=] %s\n", string(string(task.Kind)[0]), task.Name, task.Priority, task.Desc)
			case Done:
				fmt.Printf("  - [%s] %s %v [X] %s\n", string(string(task.Kind)[0]), task.Name, task.Priority, task.Desc)
			}
		}
	}
}

func SortTasks(ts []Task) []Task {
	sort.Slice(ts, func(i, j int) bool {
		if ts[i].Status != ts[j].Status {
			return ts[i].Status < ts[j].Status
		}
		return ts[i].Priority < ts[j].Priority
	})
	return ts
}

func PrintQuotes() {
	quotes_s := Pool.Random("sentences")
	quotes_z := Pool.Random("zgxw")
	quotes_w := Pool.Random("words")
	quotes := [][]string{quotes_s, quotes_z, quotes_w}
	for _, q := range quotes {
		for index, item := range q {
			fmt.Printf("\n%d %s", index+1, item)
		}
	}
	fmt.Printf("\n")
}

/*
func init() {
	Pool.CreateTableTask()
	Pool.CreateTableSection()
	Pool.CreateTableNotice()
	Pool.InsertSection("GOODDAY", "work")
	Pool.InsertSection("CAMBRIDGE", "work")
	Pool.InsertSection("LEARN", "personal")
	Pool.InsertSection("READ", "personal")
	Pool.InsertSection("SOLVE", "personal")
	Pool.InsertNotice("Read the docs!", "note")
	Pool.InsertTask("Add terminal commands", "feature", "todo", "P1", 1)
	Pool.InsertTask("Add db methods", "feature", "doing", "P1", 1)
	Pool.InsertTask("Use Tcell", "epic", "todo", "P1", 1)
	Pool.InsertTask("Hashset algorithms", "feature", "todo", "P2", 2)
	Pool.InsertTask("Add Urban & Webster", "epic", "todo", "P2", 2)
	Pool.InsertTask("Add ncurses", "epic", "todo", "P2", 2)
	Pool.InsertTask("chi for web dev", "action", "todo", "P3", 3)
	Pool.InsertTask("lf for go", "action", "todo", "P3", 3)
	Pool.InsertTask("Suckless Articles", "action", "todo", "P3", 4)
	Pool.InsertTask("The Apology", "action", "doing", "P3", 4)
	Pool.InsertTask("传习录", "action", "doing", "P3", 4)
	Pool.InsertTask("Virtualbox Guest Additions", "action", "todo", "P3", 5)
	Pool.InsertTask("Dual Monitors", "action", "todo", "P3", 5)
}
*/
