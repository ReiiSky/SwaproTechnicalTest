package adapters

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/interfaces"
	"github.com/ReiiSky/SwaproTechnical/sources/interfaces/controllers"
	"github.com/gofiber/fiber/v3"
)

type Fiber struct {
	fiber.App
	port       int
	kernel     interfaces.IKernel
	controller controllers.Controller
}

func NewFiber(kernel interfaces.IKernel, port int) Fiber {
	return Fiber{*fiber.New(), port, kernel, controllers.NewController()}
}

func (f Fiber) Run() {
	f.registerRoute().
		Listen(fmt.Sprintf(":%d", f.port))
}

func (f *Fiber) registerRoute() *Fiber {
	f.Post("/signin", f.SignIn)
	f.Post("/employee", f.RegisterEmployee)
	f.Get("/employee", f.GetEmployeeInfo)

	// TODO: Event not implemented yet.
	f.Delete("/employee", f.DeleteEmployee)
	f.Put("/employee/superior", f.AssignSuperior)
	f.Get("/position", f.GetPositionInformation)
	f.Post("/position", f.ApplyPosition)
	f.Put("/position", f.ChangePositionName)
	f.Delete("/position", f.DeletePosition)
	f.Get("/department", f.GetDepartmentInformation)
	f.Put("/department", f.ChangeDepartmentName)
	f.Delete("/department", f.DeleteDepartment)
	f.Post("/checkin", f.CheckIn)
	f.Put("/checkout", f.CheckOut)
	f.Delete("/attendance", f.DeleteAttendance)
	f.Get("/attendance/location", f.GetLocationAttendance)
	f.Put("/attendance/location", f.ChangeLocationName)
	f.Delete("/location", f.DeleteLocation)

	return f
}

func parseBearer(auth string) *string {
	if strings.Index(auth, "Bearer") != 0 {
		return nil
	}

	auths := strings.Split(auth, " ")

	if len(auths) != 2 {
		return nil
	}

	return &auths[1]
}

func (f Fiber) parse(ctx fiber.Ctx) controllers.ControllerPayload {
	payload := controllers.ControllerPayload{}
	headers := ctx.GetReqHeaders()
	queries := ctx.Queries()

	if authorization := headers["Authorization"]; len(authorization) > 0 {
		payload.Authtoken = parseBearer(authorization[0])
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

func (f *Fiber) SignIn(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	output, errCode := f.controller.Login(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.created(ctx, output)
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

func (f *Fiber) ChangePositionName(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.ChangePositionName(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) DeletePosition(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.DeletePosition(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) GetDepartmentInformation(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	output, errCode := f.controller.GetDepartmentInformation(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, output)
}

func (f *Fiber) ChangeDepartmentName(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.ChangeDepartmentName(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) DeleteDepartment(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.DeleteDepartment(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) CheckIn(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.CheckIn(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) CheckOut(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.CheckOut(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) DeleteAttendance(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.DeleteAttendance(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) GetLocationAttendance(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	output, errCode := f.controller.GetLocationAttendance(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, output)
}

func (f *Fiber) ChangeLocationName(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.ChangeLocationName(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}

func (f *Fiber) DeleteLocation(ctx fiber.Ctx) error {
	process := f.kernel.NewProcess()
	defer process.Close()

	requestPayload := f.parse(ctx)
	errCode := f.controller.DeleteLocation(process, requestPayload)

	if errCode != nil {
		return f.apply(ctx, *errCode)
	}

	return f.ok(ctx, nil)
}
