package main

import "fmt"

// 1. Define a custom type based on an integer
type Status int

// 2. Declare the values in a const block using iota
const (
	StatusUnknown   Status = iota // 0 (Idiomatic zero-value default)
	StatusPending                 // 1
	StatusActive                  // 2
	StatusCompleted               // 3
)

func main() {
	var current Status = StatusPending
	fmt.Println(current) // Prints: 1
	result := getworkDayStatus("sunday")

	fmt.Println(result)

}

func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "Unknown"
	case StatusPending:
		return "Pending"
	case StatusActive:
		return "Active"
	case StatusCompleted:
		return "Completed"
	default:
		return "Unknown"
	}
}

func getworkDayStatus(day string) string {
	switch day {
	case "sunday", "monday", "tuesday", "wednesday":
		return "Office is open"
	case "thursday":
		return "office is close"
	default:
		return "invalid day"
	}
}
