package inmemorydb

import (
	"context"

	"github.com/Lefthander/otus-go-calendar/internal/domain/entities"
	"github.com/Lefthander/otus-go-calendar/internal/domain/errors"
)

// InMemoryStore is a slice of CalendarEvents struct items
type InMemoryStore struct {
	Store map[uint64]entities.CalendarEvent
}

func newInMemoryStore(numberOfItems uint) *InMemoryStore {
	ims := &InMemoryStore{Store: make(map[uint64]entities.CalendarEvent, numberOfItems)}
	return ims
}

// SaveEvent stores the CalendarEvent item in slice
func (ims *InMemoryStore) SaveEvent(ctx context.Context, event *entities.CalendarEvent) error {
	if _, ok := ims.Store[event.ID]; ok {
		return errors.ErrCalendarUnableToSaveData
	}
	ims.Store[event.ID] = *event
	return nil
}

func (ims *InMemoryStore) GetEvent(ctx context.Context, eventID uint64) (*entities.CalendarEvent, error) {
	if event, ok := ims.Store[eventID]; ok {
		return &event, nil
	}
	return nil, errors.ErrCalendarNoSuchEvent
}

func (ims *InMemoryStore) DeleteEvent(ctx context.Context, eventID uint64) error {
	return nil
}
