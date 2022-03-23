// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// https://pkg.go.dev/os/signal?tab=licenses
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("Got args: %v\n", os.Args[1:])
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go spinner()

	// Block until a signal is received.
	s := <-c
	fmt.Printf("\rGot signal: %s\n", s)
	os.Exit(0)
}

func spinner() {
	for {
		for _, r :=range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
