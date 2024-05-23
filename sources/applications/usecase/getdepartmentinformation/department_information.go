package getdepartmentinformation

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/services/auth"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	domainErr "github.com/ReiiSky/SwaproTechnical/sources/domains/errors"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
)

type Usecase struct{}

var errMaps = usecase.NewErrorMapper().
	Add(usecase.ErrCodeNotFound, domainErr.EmployeeNotExist{}).
	Add(usecase.ErrCodeNotFound, domainErr.PositionOrDepartmentNotExist{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
) (DepartmentInformationOutput, *usecase.ErrorWithCode) {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return DepartmentInformationOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	department := employee.Department()

	if department == nil {
		return DepartmentInformationOutput{}, errMaps.Map(domainErr.PositionOrDepartmentNotExist{})
	}

	return DepartmentInformationOutput{
		ID:            objects.GetNumberIdentifier(department.ID()),
		Name:          department.Name(),
		EmployeeCount: department.Info().EmployeeCount,
		PositionCount: department.Info().PositionCount,
		CreatedAt:     department.Changelog().CreatedAt().ToISOUTC(),
	}, nil
}
