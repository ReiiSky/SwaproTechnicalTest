package register

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	domainErr "github.com/ReiiSky/SwaproTechnical/sources/domains/errors"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
)

type Usecase struct{}

func NewRegisterUsecase() Usecase {
	currentUsecase := Usecase{}

	return currentUsecase
}

var errMaps = usecase.NewErrorMapper().
	Add(usecase.ErrCodeConflict, domainErr.EmployeeIsExist{})

func (u Usecase) Execute(
	process applications.Process,
	input EmployeeRegisterInput,
) *usecase.ErrorWithCode {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByName{Name: input.Name})

	_, ok := aggr.(*domains.Employee)

	if ok {
		return errMaps.Map(domainErr.EmployeeIsExist{})
	}

	employee := repositories.Employee().
		GetOne(specifications.GetEmpty{
			Name:     input.Name,
			Password: process.Services().Hasher().Hash(input.Password),
		}).(*domains.Employee)

	employee.Register()
	repositories.Save(employee)

	return nil
}
