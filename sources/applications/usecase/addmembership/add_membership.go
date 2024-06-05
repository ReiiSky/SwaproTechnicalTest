package addmembership

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
	Add(usecase.ErrCodeConflict, domainErr.AlreadEmployee{}).
	Add(usecase.ErrCodeNotFound, domainErr.PositionOrDepartmentNotExist{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
	input AddMembershipInput,
) *usecase.ErrorWithCode {
	var (
		aggr         applications.Aggregate
		repositories = process.Repositories()
	)

	aggr = repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return errMaps.Map(domainErr.EmployeeNotExist{})
	}

	hashedPasword := process.Services().Hasher().Hash(input.Password)
	err := employee.AddMembership(input.Name, hashedPasword, input.Address)

	if err != nil {
		return errMaps.Map(err)
	}

	repositories.Save(employee)

	return nil
}
