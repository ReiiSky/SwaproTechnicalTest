package controllers

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/addmembership"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/applytoposition"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/assignsupervisor"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/changedepartmentname"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/changelocationname"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/changepositionname"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/checkin"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/checkout"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deleteattendance"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deletedepartment"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deleteemployee"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deletelocation"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/deleteposition"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getdepartmentinformation"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getemployeeinfo"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getlocationattendances"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getmembership"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/getpositioninformation"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/login"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase/register"
)

type Controller struct {
	registerUsecase               register.Usecase
	loginUsecase                  login.Usecase
	getemployeeInfoUsecase        getemployeeinfo.Usecase
	deleteemployeeUsecase         deleteemployee.Usecase
	assignsuperior                assignsupervisor.Usecase
	getpositioninformationUsecase getpositioninformation.Usecase
	applytopositionUsecase        applytoposition.Usecase
	changepositionUsecase         changepositionname.Usecase
	deletepositionUsecase         deleteposition.Usecase
	getdepartmentinformation      getdepartmentinformation.Usecase
	changedepartmentname          changedepartmentname.Usecase
	deletedepartmentUsecase       deletedepartment.Usecase
	checkinUsecase                checkin.Usecase
	checkoutUsecase               checkout.Usecase
	deleteattendance              deleteattendance.Usecase
	getlocationattendances        getlocationattendances.Usecase
	changelocationname            changelocationname.Usecase
	deletelocationUsecase         deletelocation.Usecase
	addmembershipUsecase          addmembership.Usecase
	getmembershipUsecase          getmembership.Usecase
}

func NewController() Controller {
	return Controller{
		register.Usecase{},
		login.Usecase{},
		getemployeeinfo.Usecase{},
		deleteemployee.Usecase{},
		assignsupervisor.Usecase{},
		getpositioninformation.Usecase{},
		applytoposition.Usecase{},
		changepositionname.Usecase{},
		deleteposition.Usecase{},
		getdepartmentinformation.Usecase{},
		changedepartmentname.Usecase{},
		deletedepartment.Usecase{},
		checkin.Usecase{},
		checkout.Usecase{},
		deleteattendance.Usecase{},
		getlocationattendances.Usecase{},
		changelocationname.Usecase{},
		deletelocation.Usecase{},
		addmembership.Usecase{},
		getmembership.Usecase{},
	}
}

type ControllerPayload struct {
	Authtoken  *string
	Query      map[string]string
	BodyString *string
}

func (c Controller) Login(process applications.Process, payload ControllerPayload) (login.LoginOutput, *usecase.ErrorWithCode) {
	if payload.BodyString == nil {
		return login.LoginOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: nil,
		}
	}

	input, err := login.NewEmployeeLoginInput(*payload.BodyString)

	if err != nil {
		return login.LoginOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.loginUsecase.Execute(process, input)
}

func (c Controller) RegisterEmployee(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

func (c Controller) GetEmployeeInfo(process applications.Process, payload ControllerPayload) (getemployeeinfo.EmployeeInfoOutput, *usecase.ErrorWithCode) {
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

func (c Controller) DeleteEmployee(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

func (c Controller) AssignSuperior(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

func (c Controller) GetPositionInformation(process applications.Process, payload ControllerPayload) (getpositioninformation.PositionInformationOutput, *usecase.ErrorWithCode) {
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

func (c Controller) ApplyPosition(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

func (c Controller) ChangePositionName(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	input, err := changepositionname.NewChangePositionInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	errCode := c.changepositionUsecase.Execute(process, authPayload, input)
	return errCode
}

func (c Controller) DeletePosition(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	errCode := c.deletepositionUsecase.Execute(process, authPayload)
	return errCode
}

func (c Controller) GetDepartmentInformation(process applications.Process, payload ControllerPayload) (getdepartmentinformation.DepartmentInformationOutput, *usecase.ErrorWithCode) {
	if payload.Authtoken == nil {
		return getdepartmentinformation.DepartmentInformationOutput{}, &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return getdepartmentinformation.DepartmentInformationOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	return c.getdepartmentinformation.Execute(process, authPayload)
}

func (c Controller) ChangeDepartmentName(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	input, err := changedepartmentname.NewChangeDepartmentInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.changedepartmentname.Execute(process, authPayload, input)
}

func (c Controller) DeleteDepartment(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	return c.deletedepartmentUsecase.Execute(process, authPayload)
}

func (c Controller) CheckIn(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	input, err := checkin.NewCheckInInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.checkinUsecase.Execute(process, authPayload, input)
}

func (c Controller) CheckOut(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	return c.checkoutUsecase.Execute(process, authPayload)
}

func (c Controller) DeleteAttendance(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	return c.deleteattendance.Execute(process, authPayload)
}

func (c Controller) GetLocationAttendance(process applications.Process, payload ControllerPayload) ([]getlocationattendances.LocationAttendanceOutput, *usecase.ErrorWithCode) {
	if payload.Authtoken == nil {
		return []getlocationattendances.LocationAttendanceOutput{}, &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return []getlocationattendances.LocationAttendanceOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	return c.getlocationattendances.Execute(process, authPayload)
}

func (c Controller) ChangeLocationName(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	input, err := changelocationname.NewChangeLocationNameInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.changelocationname.Execute(process, authPayload, input)
}

func (c Controller) DeleteLocation(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	input, err := deletelocation.NewDeleteLocationInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.deletelocationUsecase.Execute(process, authPayload, input)
}

func (c Controller) AddMembership(process applications.Process, payload ControllerPayload) *usecase.ErrorWithCode {
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

	input, err := addmembership.NewAddMembershipInput(*payload.BodyString)

	if err != nil {
		return &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeInvalidRequest,
			ErrInstance: err,
		}
	}

	return c.addmembershipUsecase.Execute(process, authPayload, input)
}

func (c Controller) GetMembership(process applications.Process, payload ControllerPayload) (getmembership.MembershipOutput, *usecase.ErrorWithCode) {
	if payload.Authtoken == nil {
		return getmembership.MembershipOutput{}, &usecase.ErrorWithCode{
			ErrCode: usecase.ErrCodeUnauthorized,
		}
	}

	authPayload, err := process.Services().
		Auth().Decode(*payload.Authtoken)

	if err != nil {
		return getmembership.MembershipOutput{}, &usecase.ErrorWithCode{
			ErrCode:     usecase.ErrCodeUnauthorized,
			ErrInstance: err,
		}
	}

	return c.getmembershipUsecase.Execute(process, authPayload)
}
