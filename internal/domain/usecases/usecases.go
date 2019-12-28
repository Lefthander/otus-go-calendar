package usecases

import (
	"context"
	"time"

	"github.com/Lefthander/otus-go-calendar/internal/domain/entities"
	"github.com/Lefthander/otus-go-calendar/internal/domain/interfaces"
	"github.com/Lefthander/otus-go-calendar/internal/domain/utils"
)

// CalendarUseCases is a struct :)
type CalendarUseCases struct {
	CalendarStorage interfaces.EventKeeper
}

// CreateCalendarEvent ...
func (usecase *CalendarUseCases) CreateCalendarEvent(ctx context.Context, user uint64, title, note string, repeated bool, repeat, start, end time.Time) (*entities.CalendarEvent, error) {
	event := &entities.CalendarEvent{
		ID:        utils.NewID(),
		Title:     title,
		Note:      note,
		UserID:    user,
		BeginTime: start,
		EndTime:   end,
	}
	if err := event.Validate(); err != nil {
		return nil, err
	}
	err := usecase.CalendarStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// GetEvent returns event from the store by id
func (usecase *CalendarUseCases) GetEvent(ctx context.Context, eventID uint64) (*entities.CalendarEvent, error) {
	return nil, nil
}
