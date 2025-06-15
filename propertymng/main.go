package main

import (
	"fmt"
	"time"
)

type State string

const (
	Idle     State = "idle"
	Running  State = "running"
	Finished State = "finished"
)

func main() {
	state := Idle

	for {
		switch state {
		case Idle:
			fmt.Println("Initializing system...")
			state = startTask()
		case Running:
			fmt.Println("System is running...")
			if taskDone() {
				state = Finished
			} else {
				time.Sleep(time.Second)
			}
		case Finished:
			fmt.Println("System is shutting down.")
			cleanup()
			return
		default:
			fmt.Println("Unknown state!")
			return
		}
	}
}

// Simulate starting a task
func startTask() State {
	fmt.Println("Starting task...")
	return Running
}

// Simulate checking if task is done
func taskDone() bool {
	return time.Now().Second()%5 == 0 // simulate occasional finish
}

// Clean up resources
func cleanup() {
	fmt.Println("Cleaning up...")
}
