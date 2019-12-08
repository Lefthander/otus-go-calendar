package entities

import (
	"time"
)

// Event represents a model of the Calendar Event
type Event struct {
	ID        uint64    // ID of Event
	Title     string    // Title of event
	Note      string    // Event notes
	UserID    uint64    // ID of user
	BeginTime time.Time // Event start time
	EndTime   time.Time // Event end time
}
