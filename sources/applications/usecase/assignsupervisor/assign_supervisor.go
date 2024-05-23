package assignsupervisor

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
	Add(usecase.ErrCodeNotFound, domainErr.DepartmentNotExist{})

func (u Usecase) Execute(
	process applications.Process,
	authPayload auth.AuthPayload,
	input AssignSupervisorInput,
) *usecase.ErrorWithCode {
	var (
		aggr         applications.Aggregate
		repositories = process.Repositories()
	)

	aggr = repositories.Employee().
		GetOne(specifications.GetByID{ID: authPayload.EmployeeID})

	subordinate, ok := aggr.(*domains.Employee)

	if !ok {
		return errMaps.Map(domainErr.EmployeeNotExist{})
	}

	aggr = repositories.Employee().
		GetOne(specifications.GetByID{ID: input.SuperiorID})

	superior, ok := aggr.(*domains.Employee)

	if !ok {
		return errMaps.Map(domainErr.EmployeeNotExist{})
	}

	err := subordinate.AssignSuperior(superior, domains.PositionParam{
		Name: input.NewPositionName,
	})

	if err != nil {
		return errMaps.Map(err)
	}

	repositories.Save(subordinate)

	return nil
}
