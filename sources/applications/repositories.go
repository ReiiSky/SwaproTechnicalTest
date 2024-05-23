package applications

import "github.com/ReiiSky/SwaproTechnical/sources/domains"

type Aggregate interface {
	Events() []domains.EventIDDecorator
}

type QueryRepository interface {
	GetOne(domains.ISpecification) Aggregate
	Get(domains.ISpecification) []Aggregate
}

type CommandRepository interface {
	Save(Aggregate)
}
