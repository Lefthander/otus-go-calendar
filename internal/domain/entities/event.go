package event

import (
	"time"
)

// CalendarEvent is a struct to store the Calendar Events
type CalendarEvent struct {
	ID        string
	Title     string
	Note      string
	UserID    string
	BeginTime time.Time
	EndTime   time.Time
}
