package controllers

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deleteemployee"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getemployeeinfo"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/register"
)

type Controller struct {
	registerUsecase        register.Usecase
	getemployeeInfoUsecase getemployeeinfo.Usecase
	deleteemployeeUsecase  deleteemployee.Usecase
}

func NewController() Controller {
	return Controller{
		register.Usecase{},
		getemployeeinfo.Usecase{},
		deleteemployee.Usecase{},
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
) (getemployeeinfo.EmployeeRegisterOutput, *usecase.ErrorWithCode) {
	if payload.Authtoken == nil {
		return getemployeeinfo.EmployeeRegisterOutput{}, &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return getemployeeinfo.EmployeeRegisterOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	output, errCode := c.getemployeeInfoUsecase.Execute(process, authPayload)

	if errCode != nil {
		return getemployeeinfo.EmployeeRegisterOutput{}, errCode
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
