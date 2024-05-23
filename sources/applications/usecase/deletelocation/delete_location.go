package deletelocation

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
	Add(usecase.ErrCodeNotFound, domainErr.LocationNotExist{}).
	Add(usecase.ErrCodeNotFound, domainErr.AttendanceNotFound{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
	input DeleteLocationInput,
) *usecase.ErrorWithCode {
	var (
		aggr         applications.Aggregate
		repositories = process.Repositories()
	)

	aggr = repositories.Employee().
		GetOne(specifications.GetByID{
			ID:              authPayload.EmployeeID,
			AttendanceLimit: 100,
		})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return errMaps.Map(domainErr.EmployeeNotExist{})
	}

	att := employee.GetAttendanceByID(input.AttendanceID)

	if att == nil {
		return errMaps.Map(domainErr.AttendanceNotFound{})
	}

	err := employee.DeleteLocationByAttendance(att)

	if err != nil {
		return errMaps.Map(err)
	}

	repositories.Save(employee)

	return nil
}
