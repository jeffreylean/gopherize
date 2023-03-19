package main

import "time"

// I AM NOT DONE
func someProcess() {
	time.Sleep(time.Second * 3)
}

func main() {
	go someProcess()
}
