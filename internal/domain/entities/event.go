package entities

import (
	"time"
)

// CalendarEvent represents a model of the Calendar Event
type CalendarEvent struct {
	ID        uint64    // ID of Event
	Title     string    // Title of event
	Note      string    // Event notes
	UserID    uint64    // ID of user
	BeginTime time.Time // Event start time
	EndTime   time.Time // Event end time
}

// Validate checks the consistancy on event paramters if beginTime is greater than EndTime returns false
func (ce *CalendarEvent) Validate() bool {
	if ce.BeginTime.After(ce.EndTime) {
		return false
	}
	return true
}
