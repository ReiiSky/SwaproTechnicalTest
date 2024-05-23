package controllers

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/applytoposition"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/assignsupervisor"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deleteemployee"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getemployeeinfo"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getpositioninformation"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/register"
)

type Controller struct {
	registerUsecase               register.Usecase
	getemployeeInfoUsecase        getemployeeinfo.Usecase
	deleteemployeeUsecase         deleteemployee.Usecase
	assignsuperior                assignsupervisor.Usecase
	getpositioninformationUsecase getpositioninformation.Usecase
	applytopositionUsecase        applytoposition.Usecase
}

func NewController() Controller {
	return Controller{
		register.Usecase{},
		getemployeeinfo.Usecase{},
		deleteemployee.Usecase{},
		assignsupervisor.Usecase{},
		getpositioninformation.Usecase{},
		applytoposition.Usecase{},
	}
}

type ControllerPayload struct {
	Authtoken  *string
	Query      map[string]string
	BodyString *string
}

func (c Controller) RegisterEmployee(
	process applications.Process,
	payload ControllerPayload,
) *usecase.ErrorWithCode {
	if payload.BodyString == nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: nil,
		}
	}

	input, err := register.NewEmployeeRegisterInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.registerUsecase.Execute(process, input)
}

func (c Controller) GetEmployeeInfo(
	process applications.Process,
	payload ControllerPayload,
) (getemployeeinfo.EmployeeInfoOutput, *usecase.ErrorWithCode) {
	if payload.Authtoken == nil {
		return getemployeeinfo.EmployeeInfoOutput{}, &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return getemployeeinfo.EmployeeInfoOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	output, errCode := c.getemployeeInfoUsecase.Execute(process, authPayload)

	if errCode != nil {
		return getemployeeinfo.EmployeeInfoOutput{}, errCode
	}

	return output, nil
}

func (c Controller) DeleteEmployee(
	process applications.Process,
	payload ControllerPayload,
) *usecase.ErrorWithCode {
	if payload.Authtoken == nil {
		return &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	errCode := c.deleteemployeeUsecase.Execute(process, authPayload)

	if errCode != nil {
		return errCode
	}

	return nil
}

func (c Controller) AssignSuperior(
	process applications.Process,
	payload ControllerPayload,
) *usecase.ErrorWithCode {
	if payload.Authtoken == nil {
		return &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	input, err := assignsupervisor.NewAssignSupervisorInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	errCode := c.assignsuperior.Execute(process, authPayload, input)

	if errCode != nil {
		return errCode
	}

	return nil
}

func (c Controller) GetPositionInformation(
	process applications.Process,
	payload ControllerPayload,
) (getpositioninformation.PositionInformationOutput, *usecase.ErrorWithCode) {
	if payload.Authtoken == nil {
		return getpositioninformation.PositionInformationOutput{}, &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return getpositioninformation.PositionInformationOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	output, errCode := c.getpositioninformationUsecase.Execute(process, authPayload)

	if errCode != nil {
		return getpositioninformation.PositionInformationOutput{}, errCode
	}

	return output, nil
}

func (c Controller) ApplyPosition(
	process applications.Process,
	payload ControllerPayload,
) *usecase.ErrorWithCode {
	if payload.Authtoken == nil {
		return &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	input, err := applytoposition.NewApplyToPositionInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	errCode := c.applytopositionUsecase.Execute(process, authPayload, input)

	if errCode != nil {
		return errCode
	}

	return nil
}
