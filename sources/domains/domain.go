package domains

const (
	ScopeEmployee = "employee"
)

type EventIDDecorator[T IEvent] struct {
	// event ID will be used for retrieve returned value
	// such as generated id, or returned error.
	// after event published to command repository.
	id  int
	top T
}

func (dec EventIDDecorator[T]) ID() int {
	return dec.id
}

func (dec EventIDDecorator[T]) Top() T {
	return dec.top
}

type IEvent interface {
	Eventname() string
}

type AggregateContext struct {
	events         []EventIDDecorator[IEvent]
	currentEventID int
}

func NewContext() AggregateContext {
	return AggregateContext{
		events:         make([]EventIDDecorator[IEvent], 0),
		currentEventID: 1,
	}
}

func (aggr *AggregateContext) addEvent(event IEvent) EventIDDecorator[IEvent] {
	if aggr.events == nil {
		aggr.events = make([]EventIDDecorator[IEvent], 0)
	}

	eventWithID := EventIDDecorator[IEvent]{id: aggr.currentEventID, top: event}
	aggr.events = append(aggr.events, eventWithID)
	aggr.currentEventID++

	return eventWithID
}

func (aggr AggregateContext) Events() []EventIDDecorator[IEvent] {
	return aggr.events
}
