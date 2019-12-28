package errors

// SentitelError is an Application Level Error
type SentitelError string

func (se SentitelError) Error() string {
	return string(se)
}

var (
	// Calendar Errors
	ErrCalendarEventOverlaping   = SentitelError("Another event is already exists on that time")
	ErrCalendarInvalidDateFormat = SentitelError("Invalid date format is used")

	// Storage Errors
	ErrCalendarUnableToSaveData = SentitelError("Unable to save event data")
	ErrCalendarUnableToReadData = SentitelError("Unable to read event data")
	ErrCalendarNoSuchEvent      = SentitelError("There is no such event in store")
)
