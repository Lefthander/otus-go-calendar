package storage

import (
	"github.com/lefthander/otus-go-calendar/internal/domain/entities"
)

// EventKeeper interface with the storage backend
type EventKeeper interface {
	SaveEvent() error
	GetEvent() (*entities.CalendarEvent, error)
	GetAllEvents() ([]entities.CalendarEvent, error)
}
