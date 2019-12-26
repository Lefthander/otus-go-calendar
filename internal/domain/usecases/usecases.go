package usecases
import (
	"github.com/lefthander/otus-go-calendar/internal/domain/interfaces"
	"github.com/lefthander/otus-go-calendar/internal/domain/entities"
	"github.com/lefthander/otus-go-calendar/internal/domain/utils"

)
// CalendarUseCases is a struct :)
type CalendarUseCases struct {
	CalendarStorage interfaces.EventKeeper
}

func (ces *CalendarUseCases) CreateCalendarEvent(ctx context.Context, user uint64, title, note string, \ 
	                                             repeated bool, repeat, start, end *time.Time) (*entities.CalendarEvent, error){
     event:= &entities.CalendarEvent{
		 ID: utils.NewID(),
		 Title: title,
		 Note: note,
		 isRepeated: repeated,
		 RepeatTime: rt,
		 UserID: user,
		 BeginTime: start,
		 EndTime: end,
	 }
	 err:= ces.CalendarStorage.SaveEvent(ctx,event)
	 if err != nil {
		 return nil, err
	 }
	 return event,nil
}