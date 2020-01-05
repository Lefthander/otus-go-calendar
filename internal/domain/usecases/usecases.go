package usecases

import (
	"context"
	"time"

	"github.com/Lefthander/otus-go-calendar/internal/domain/entities"
	"github.com/Lefthander/otus-go-calendar/internal/domain/errors"
	"github.com/Lefthander/otus-go-calendar/internal/domain/interfaces"
	"github.com/Lefthander/otus-go-calendar/internal/domain/utils"
)

// CalendarUseCases is a struct :)
type CalendarUseCases struct {
	CalendarStorage interfaces.EventKeeper
}

// CreateEvent ...
func (usecase *CalendarUseCases) CreateEvent(ctx context.Context, user uint64, title, note string, start, end time.Time) (*entities.CalendarEvent, error) {
	event := &entities.CalendarEvent{
		ID:        utils.NewID(),
		Title:     title,
		Note:      note,
		UserID:    user,
		BeginTime: start,
		EndTime:   end,
	}
	// Does basic validation of that begin date/time is less than end date/time
	if !event.Validate() {
		return nil, errors.ErrCalendarInvalidDateFormat
	}
	event, err := usecase.CalendarStorage.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// ReadEvent returns event from the store by event ID
func (usecase *CalendarUseCases) ReadEvent(ctx context.Context, eventID uint64) (*entities.CalendarEvent, error) {
	// TODO:
	event, err := usecase.CalendarStorage.ReadEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// UpdateEvent updates the specific event, in case of success returns nil
func (usecase *CalendarUseCases) UpdateEvent(ctx context.Context, eventID uint64, title, note string, startTime, endTime time.Time) error {
	// TODO:
	return nil
}

// DeleteEvent deletes the event from the store, in casa of absense of requested event returns corresponding error, in case of success returns nil.
func (usecase *CalendarUseCases) DeleteEvent(ctx context.Context, eventID uint64) error {
	// TODO:
	return nil
}
