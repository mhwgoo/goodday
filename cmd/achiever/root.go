package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mhwgoo/achieve"
	"github.com/spf13/cobra"
)

var user = os.Getenv("USER")

var rootCmd = &cobra.Command{
	Use: os.Args[0],
	// Short: "A terminal tool for managing tasks",
	// Long:  "Achieve is intended to managing personal and work tasks especially for developers. It's Asana-like, but clean, not bloated, convinient.",
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

	achieve.Do()
}
