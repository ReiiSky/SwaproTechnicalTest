package getlocationattendances

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
	Add(usecase.ErrCodeNotFound, domainErr.EmployeeNotExist{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
) ([]LocationAttendanceOutput, *usecase.ErrorWithCode) {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByID{
			ID:              authPayload.EmployeeID,
			AttendanceLimit: 100,
		})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return []LocationAttendanceOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	outputs := []LocationAttendanceOutput{}
	locations := employee.UniqueAttendanceLocation()

	for _, loc := range locations {
		outputs = append(outputs, LocationAttendanceOutput{
			ID:        objects.GetNumberIdentifier(loc.ID()),
			Name:      loc.Name(),
			CreatedAt: loc.Changelog().CreatedAt().ToISOUTC(),
		})
	}

	return outputs, nil
}
