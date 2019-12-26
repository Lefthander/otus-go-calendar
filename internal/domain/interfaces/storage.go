package storage

import (
	"context"
	"time"
	"github.com/lefthander/otus-go-calendar/internal/domain/entities"
)

// EventKeeper interface with the storage backend
type EventKeeper interface {
	SaveEvent(ctx context.Context, event *entities.CalendarEvent) error
	GetEventByID(ctx context.Context, eventID uint64) (*entities.CalendarEvent, error)
	GetAllEventsByUserID(ctx context.Context, userID uint64) ([]entities.CalendarEvent, error)
	GetEventsByTimeForUser(ctx context.Context, time time.Time, userID uint64) ([]entities.CalendarEvent, error)
}
