package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "---->", i)
	}
}

func main() {
	// 1. The Synchronous way
	say("Sync")

	// 2. Go routines
	go say("Async 1")
	go say("Async 2")
	go say("Async 3")

	// Sleep for 3 sec
	time.Sleep(time.Second * 3)
}
