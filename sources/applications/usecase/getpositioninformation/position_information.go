package getpositioninformation

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
) (PositionInformationOutput, *usecase.ErrorWithCode) {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return PositionInformationOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	position := employee.Position()

	if position == nil {
		return PositionInformationOutput{}, errMaps.Map(domainErr.PositionOrDepartmentNotExist{})
	}

	return PositionInformationOutput{
		ID:             objects.GetNumberIdentifier(position.ID()),
		Name:           position.Name(),
		DepartmentName: employee.Department().Name(),
		EmployeeCount:  position.Info().EmployeeCount,
		CreatedAt:      position.Changelog().CreatedAt().ToISOUTC(),
	}, nil
}
