package main

import "time"

func someProcess() {
	time.Sleep(time.Second * 3)
}

func main() {
	go someProcess()
}
