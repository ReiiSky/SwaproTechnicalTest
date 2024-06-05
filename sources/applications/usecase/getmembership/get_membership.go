package getmembership

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/services/auth"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	domainErr "github.com/ReiiSky/SwaproTechnical/sources/domains/errors"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
)

type Usecase struct{}

var errMaps = usecase.NewErrorMapper().
	Add(usecase.ErrCodeNotFound, domainErr.EmployeeNotExist{}).
	Add(usecase.ErrCodeConflict, domainErr.AlreadEmployee{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
) (MembershipOutput, *usecase.ErrorWithCode) {
	var (
		aggr         applications.Aggregate
		repositories = process.Repositories()
	)

	aggr = repositories.Employee().
		GetOne(specifications.GetByID{
			ID:             authPayload.EmployeeID,
			WithMembership: true,
		})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return MembershipOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	membership := employee.Membership()

	return MembershipOutput{
		Name:      membership.Name(),
		Address:   membership.Address(),
		IsActive:  membership.IsActive(),
		CreatedAt: membership.Changelog().CreatedAt().ToISOUTC(),
	}, nil
}
