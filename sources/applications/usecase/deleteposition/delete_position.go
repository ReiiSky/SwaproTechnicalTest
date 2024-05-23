package deleteemployee

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
	Add(usecase.ErrCodeInvalidRequest, domainErr.NotAnEmployee{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
) *usecase.ErrorWithCode {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return errMaps.Map(domainErr.EmployeeNotExist{})
	}

	err := employee.DeletePosition()

	if err != nil {
		return errMaps.Map(err)
	}

	repositories.Save(employee)

	return nil
}
