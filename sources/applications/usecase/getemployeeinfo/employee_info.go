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
) (EmployeeRegisterOutput, *usecase.ErrorWithCode) {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return EmployeeRegisterOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	info, err := employee.Info()

	if err != nil {
		return EmployeeRegisterOutput{}, errMaps.Map(err)
	}

	return EmployeeRegisterOutput{
		Name:       info.Name,
		Position:   info.PositionName,
		Department: info.DepartementName,
		CreatedAt:  info.CreatedAt.ToISOUTC(),
	}, nil
}
