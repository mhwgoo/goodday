package goodday

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mhwgoo/goodday/data"
)

func Do() {
	PrintRemember("\nREMEMBER: 常将有日思无日，莫把无时当有时")
	PrintFocus("FOCUS: goodday")
	PrintQuotes()
	PrintSections(sections)
}

func PrintSections(secs []section) {
	for _, sec := range secs {
		fmt.Printf("\n[%s] %s\n", string(string(sec.kind)[0]), strings.ToUpper(sec.name))
		for _, task := range SortTasks(sec.tasks) {
			if task.priority == p1 {
				task.priority = level("***")
			} else if task.priority == p2 {
				task.priority = level("*")
			} else {
				task.priority = level("")
			}
			switch task.status {
			case todo:
				fmt.Printf("  - [%s] %s %v [ ] %s\n", string(string(task.kind)[0]), task.name, task.priority, task.desc)
			case doing:
				fmt.Printf("  - [%s] %s %v [=] %s\n", string(string(task.kind)[0]), task.name, task.priority, task.desc)
			case done:
				fmt.Printf("  - [%s] %s %v [X] %s\n", string(string(task.kind)[0]), task.name, task.priority, task.desc)
			}
		}
	}
}

func PrintQuotes() {
	DB, _ := sql.Open("sqlite3", "./data/goodday.db")
	Pool := data.NewPool(DB)
	quotes_s := Pool.Random("sentences")
	for index, quote := range quotes_s {
		fmt.Printf("\n%d %s", index+1, quote)
	}
	quotes_z := Pool.Random("zgxw")
	for index, quote := range quotes_z {
		fmt.Printf("\n%d %s", index+1, quote)
	}
	quotes_w := Pool.Random("words")
	for index, quote := range quotes_w {
		fmt.Printf("\n%d %s", index+1, quote)
	}
	Pool.Add("words", "up in the air")
	fmt.Printf("\n")
}

func PrintRemember(text string) {
	newr := Remember{Content: text, Date: time.Now()}
	fmt.Println(newr.Content)
}

func PrintFocus(text string) {
	newf := Focus{Content: text, Date: time.Now()}
	fmt.Println(newf.Content)
}

func SortTasks(ts []task) []task {
	sort.Slice(ts, func(i, j int) bool {
		if ts[i].status != ts[j].status {
			return ts[i].status < ts[j].status
		}
		return ts[i].priority < ts[j].priority
	})
	return ts
}
