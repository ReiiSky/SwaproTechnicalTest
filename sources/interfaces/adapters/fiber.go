package adapters

import (
	"errors"
	"net/http"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/interfaces"
	"github.com/ReiiSky/SwaproTechnical/sources/interfaces/controllers"
	"github.com/gofiber/fiber/v3"
)

type Fiber struct {
	kernel     interfaces.IKernel
	controller controllers.Controller
}

func NewFiber(kernel interfaces.IKernel) Fiber {
	return Fiber{kernel, controllers.NewController()}
}

func (f Fiber) parse(ctx fiber.Ctx) controllers.ControllerPayload {
	payload := controllers.ControllerPayload{}
	headers := ctx.GetReqHeaders()
	queries := ctx.Queries()

	if authorization := headers["Authorization"]; len(authorization) > 0 {
		payload.Authtoken = &authorization[0]
	}

	payload.Query = queries

	if ctx.Method() != "GET" {
		body := string(ctx.Body())
		payload.BodyString = &body
	}

	return payload
}

func (f Fiber) ok(ctx fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(interfaces.NewOKResponse(data))
}

func (f Fiber) created(ctx fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(interfaces.NewOKResponse(data))
}

func (f Fiber) apply(ctx fiber.Ctx, errCode usecase.ErrorWithCode) error {
	switch errCode.ErrCode {
	case usecase.ErrCodeInvalidRequest:
		return ctx.Status(http.StatusBadRequest).
			JSON(interfaces.NewFailedResponse(nil, errCode.ErrInstance))
	case usecase.ErrCodeUnauthorized:
		return ctx.Status(http.StatusUnauthorized).
			JSON(interfaces.NewFailedResponse(nil, errCode.ErrInstance))
	case usecase.ErrCodeNotFound:
		return ctx.Status(http.StatusNotFound).
			JSON(interfaces.NewFailedResponse(nil, errCode.ErrInstance))
	case usecase.ErrCodeConflict:
		return ctx.Status(http.StatusConflict).
			JSON(interfaces.NewFailedResponse(nil, errCode.ErrInstance))
	}

	return ctx.Status(http.StatusInternalServerError).
		JSON(interfaces.NewFailedResponse(nil, errors.New("internal server error")))
}

func (f *Fiber) RegisterEmployee(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.RegisterEmployee(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.created(ctx, nil)
}

func (f *Fiber) GetEmployeeInfo(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	output, errCode := f.controller.GetEmployeeInfo(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, output)
}

func (f *Fiber) DeleteEmployee(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.DeleteEmployee(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) AssignSuperior(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.AssignSuperior(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) GetPositionInformation(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	output, errCode := f.controller.GetPositionInformation(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, output)
}

func (f *Fiber) ApplyPosition(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.ApplyPosition(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}
