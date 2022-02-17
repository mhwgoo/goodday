package achieve

import (
	"fmt"
	"strings"
)

func Do() {
	PrintProjects(projects)
}

func PrintProjects(pjs []project) {
	for _, pj := range pjs {
		fmt.Printf("\n[%s] %s\n", string(string(pj.kind)[0]), strings.ToUpper(pj.name))
		for _, task := range pj.tasks {
			if task.priority == HIGH {
				task.priority = level("***")
			} else if task.priority == MIDDLE {
				task.priority = level("*")
			} else {
				task.priority = level("")
			}
			switch task.status {
			case todo:
				fmt.Printf("  - [%s] %s %v [ ]\n", string(string(task.kind)[0]), task.name, task.priority)
			case doing:
				fmt.Printf("  - [%s] %s %v [=]\n", string(string(task.kind)[0]), task.name, task.priority)
			case done:
				fmt.Printf("  - [%s] %s %v [X]\n", string(string(task.kind)[0]), task.name, task.priority)
			}
		}
	}
}
