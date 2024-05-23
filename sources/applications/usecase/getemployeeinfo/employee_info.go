package getemployeeinfo

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
	Add(usecase.ErrCodeNotFound, domainErr.EmployeeNotExist{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
) (EmployeeInfoOutput, *usecase.ErrorWithCode) {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return EmployeeInfoOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	info, err := employee.Info()

	if err != nil {
		return EmployeeInfoOutput{}, errMaps.Map(err)
	}

	return EmployeeInfoOutput{
		Name:       info.Name,
		Position:   info.PositionName,
		Department: info.DepartementName,
		CreatedAt:  info.CreatedAt.ToISOUTC(),
	}, nil
}
