package achiever

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mhwgoo/achiever/data"
)

func Do() {
	PrintRemember("\nREMEMBER: Read the Docs & Posts!")
	PrintFocus("FOCUS: Learning Gorm")
	PrintQuotes()
	PrintProjects(projects)
}

func PrintProjects(pjs []project) {
	for _, pj := range pjs {
		fmt.Printf("\n[%s] %s\n", string(string(pj.kind)[0]), strings.ToUpper(pj.name))
		for _, task := range pj.tasks {
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
	DB, _ := sql.Open("sqlite3", "./data/achiever.db")
	// data.CreateTableS(DB)
	Pool := data.NewPool(DB)
	quotes_s := Pool.GetS()
	for index, quote := range quotes_s {
		fmt.Printf("\n%d %s", index+1, quote.Text)
	}
	quotes_z := Pool.GetZGXW()
	for index, quote := range quotes_z {
		fmt.Printf("\n%d %s", index+1, quote.Text)
	}
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
