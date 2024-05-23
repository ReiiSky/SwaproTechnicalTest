package domains

const (
	ScopeEmployee = "employee"
)

type EventIDDecorator struct {
	// event ID will be used for retrieve returned value
	// such as generated id, or returned error.
	// after event published to command repository.
	id  int
	top IEvent
}

func (dec EventIDDecorator) ID() int {
	return dec.id
}

func (dec EventIDDecorator) Top() IEvent {
	return dec.top
}

type IEvent interface {
	Eventname() string
}

type AggregateContext struct {
	events         []EventIDDecorator
	currentEventID int
}

func NewContext() AggregateContext {
	return AggregateContext{
		events:         make([]EventIDDecorator, 0),
		currentEventID: 1,
	}
}

func (aggr *AggregateContext) addEvent(event IEvent) EventIDDecorator {
	if aggr.events == nil {
		aggr.events = make([]EventIDDecorator, 0)
	}

	eventWithID := EventIDDecorator{id: aggr.currentEventID, top: event}
	aggr.events = append(aggr.events, eventWithID)
	aggr.currentEventID++

	return eventWithID
}

func (aggr AggregateContext) Events() []EventIDDecorator {
	return aggr.events
}
