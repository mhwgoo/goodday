package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer un(trace("Goodday"))
	exitOnError(rootCmd.Execute())
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func trace(s string) (string, time.Time) {
	fmt.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println("\nEND:", s, "ElapsedTime:", endTime.Sub(startTime))
}
