package interfaces

import (
	"context"
	"time"

	"github.com/Lefthander/otus-go-calendar/internal/domain/entities"
)

// EventKeeper is interface which implements basic CRUD with the storage backend
type EventKeeper interface {
	SaveEvent(ctx context.Context, event *entities.CalendarEvent) error
	GetEvent(ctx context.Context, eventID uint64) (*entities.CalendarEvent, error)
	UpdateEvent(ctx context.Context, eventID uint64, title, note string, startTime, endTime, repeatTime time.Time, isrepeat bool) error
	DeleteEvent(ctx context.Context, eventID uint64) error
	GetAllEventsByUserID(ctx context.Context, userID uint64) ([]entities.CalendarEvent, error)
	GetEventsByTimeForUser(ctx context.Context, startTime time.Time, userID uint64) ([]entities.CalendarEvent, error)
}
