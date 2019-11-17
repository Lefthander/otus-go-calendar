package event

import (
	"time"
)

// CalendarEvent is a struct to store the Calendar Events 
type CalendarEvent struct {
	Title     string
	Note      string
	User      string
	BeginTime time.Time
	EndTime   time.Time
}
