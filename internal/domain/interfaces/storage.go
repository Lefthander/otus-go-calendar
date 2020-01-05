package interfaces

import (
	"context"
	"time"

	"github.com/Lefthander/otus-go-calendar/internal/domain/entities"
)

// EventKeeper is interface which implements basic CRUD with the storage backend
type EventKeeper interface {
	CreateEvent(ctx context.Context, event *entities.CalendarEvent) (*entities.CalendarEvent, error)
	ReadEvent(ctx context.Context, eventID uint64) (*entities.CalendarEvent, error)
	UpdateEvent(ctx context.Context, eventID uint64, title, note string, startTime, endTime time.Time) error
	DeleteEvent(ctx context.Context, eventID uint64) error
	GetAllEventsByUserID(ctx context.Context, userID uint64) ([]entities.CalendarEvent, error)
	GetEventsByTimeForUserID(ctx context.Context, startTime time.Time, userID uint64) ([]entities.CalendarEvent, error)
}
