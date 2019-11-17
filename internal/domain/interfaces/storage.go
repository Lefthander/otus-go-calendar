package storage

import (
	"github.com/lefthander/otus-go-calendar/internal/domain/entities"
)

// EventKeeper interface with the storage backend
type EventKeeper interface {
	SaveEvent() error
	GetEvent() (*event.CalendarEvent, error)
}
