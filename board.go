package achieve

import (
	"fmt"
)

func Do() {
	PrintWorkspace(workspaces)
}

func PrintWorkspace(wspaces []workspace) {
	for _, ws := range wspaces {
		fmt.Printf("\nWORKSPACE: %s\n", ws.name)
		for index, project := range ws.projects {
			fmt.Printf("%d %s\n", index+1, project.name)
			for _, task := range project.tasks {
				fmt.Printf("  %s %d %s\n", task.status, task.priority, task.name)
			}
		}
	}
}
