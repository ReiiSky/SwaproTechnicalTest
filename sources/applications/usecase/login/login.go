package login

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

func NewRegisterUsecase() Usecase {
	currentUsecase := Usecase{}

	return currentUsecase
}

var errMaps = usecase.NewErrorMapper().
	Add(usecase.ErrCodeNotFound, domainErr.EmployeeNotExist{}).
	Add(usecase.ErrCodeInvalidRequest, domainErr.CredentialNotValid{})

func (u Usecase) Execute(
	process applications.Process,
	input LoginInput,
) (LoginOutput, *usecase.ErrorWithCode) {
	repositories := process.Repositories()
	aggr := repositories.Employee().
		GetOne(specifications.GetByName{Name: input.Name})

	employee, ok := aggr.(*domains.Employee)

	if !ok {
		return LoginOutput{}, errMaps.Map(domainErr.EmployeeNotExist{})
	}

	if !employee.SignInable(process.Services().Hasher().Hash(input.Password)) {
		return LoginOutput{}, errMaps.Map(domainErr.CredentialNotValid{})
	}

	employeeID := objects.GetNumberIdentifier(employee.ID())
	authToken := process.Services().Auth().Encode(auth.AuthPayload{EmployeeID: employeeID})

	return LoginOutput{authToken}, nil
}
