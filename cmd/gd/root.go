package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mhwgoo/goodday"
	"github.com/spf13/cobra"
)

var user = os.Getenv("USER")
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
		for _, Task := range SortTasks(sec.Tasks) {
			if Task.Priority == P1 {
				Task.Priority = Level("***")
			} else if Task.Priority == P2 {
				Task.Priority = Level("*")
			} else {
				Task.Priority = Level("")
			}
			switch Task.Status {
			case Todo:
				fmt.Printf("  - [%s] %s %v [ ] %s\n", string(string(Task.Kind)[0]), Task.Name, Task.Priority, Task.Desc)
			case Doing:
				fmt.Printf("  - [%s] %s %v [=] %s\n", string(string(Task.Kind)[0]), Task.Name, Task.Priority, Task.Desc)
			case Done:
				fmt.Printf("  - [%s] %s %v [X] %s\n", string(string(Task.Kind)[0]), Task.Name, Task.Priority, Task.Desc)
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
	DB, err := sql.Open("sqlite3", "/home/kate/go/src/github.com/mhwgoo/goodday/goodday.db")
	if err != nil {
		log.Println(err)
	}

	Pool := goodday.NewPool(DB)
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
	Pool.CreateTableTaskType()
	Pool.CreateTableTaskStatus()
	Pool.CreateTableTaskPriority()
	Pool.CreateTableSection()
	Pool.CreateTableSectionType()
	Pool.CreateTableSectionStatus()
	Pool.CreateTableNotice()
	Pool.CreateTableNoticeType()
}
*/
