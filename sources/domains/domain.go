package domains

const (
	ScopeEmployee = "employee"
)

type IEvent interface {
	Name() string
}

type AggregateContext struct {
	events []IEvent
}

func NewContext() AggregateContext {
	return AggregateContext{
		events: make([]IEvent, 0),
	}
}

func (aggr *AggregateContext) AddEvent(event IEvent) AggregateContext {
	if aggr.events == nil {
		aggr.events = make([]IEvent, 0)
	}

	aggr.events = append(aggr.events, event)

	return *aggr
}

func (aggr AggregateContext) Events() []IEvent {
	return aggr.events
}
