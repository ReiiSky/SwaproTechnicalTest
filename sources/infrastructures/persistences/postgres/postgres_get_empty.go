package postgres

import (
	"context"

	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
)

type GetByEmpty struct {
	specifications.GetEmpty
}

func (impl GetByEmpty) Fn() persistences.SpecImplFn {
	return func(ctx context.Context, spec domains.ISpecification) []applications.Aggregate {
		var (
			fnSpec = spec.(specifications.GetEmpty)
		)

		employee := domains.NewEmployee(0, domains.EmployeeParam{
			Name:     fnSpec.Name,
			Password: fnSpec.Password,
		})

		return []applications.Aggregate{&employee}
	}
}
