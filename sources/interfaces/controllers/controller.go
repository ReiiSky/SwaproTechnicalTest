package controllers

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/register"
)

type Controller struct {
	registerUsecase register.Usecase
}

func NewController() Controller {
	return Controller{
		register.Usecase{},
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
