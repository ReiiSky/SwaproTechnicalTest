package domains

import (
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/domains/entities"
	domainErr "github.com/ReiiSky/SwaproTechnical/sources/domains/errors"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/events"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

const (
	formatPrefixEmployeeCode = "yymm"
)

type DepartmentParam struct {
	ID   int
	Name string
	objects.ChangelogParam
}

type PositionParam struct {
	ID         int
	Name       string
	Department DepartmentParam
	objects.ChangelogParam
}

type AttendanceParam struct {
	ID         int
	EmployeeID int
	LocationID int
	AbsentIn   time.Time
	AbsentOut  *time.Time
}

type EmployeeParam struct {
	Code        string
	Position    *PositionParam
	SuperiorID  *int
	Name        string
	Password    string
	Attendances []AttendanceParam
	objects.ChangelogParam
}

type Employee struct {
	AggregateContext
	root        entities.Employee
	position    *entities.Position
	department  *entities.Department
	superiorID  *objects.Identifier[int]
	attendances []entities.Attendance

	// internal state
	isRegistered bool
	isDeleted    bool
}

func NewEmployee(id int, param EmployeeParam) Employee {
	var (
		position   *entities.Position
		department *entities.Department
		superiorID *objects.Identifier[int]
	)

	if param.Position != nil {
		departmentParam := param.Position.Department
		d := entities.NewDepartment(
			departmentParam.ID,
			departmentParam.Name,
			departmentParam.ChangelogParam,
		)

		p := entities.NewPosition(
			param.Position.ID,
			param.Position.Department.ID,
			param.Position.Name,
			param.Position.ChangelogParam,
		)

		position = &p
		department = &d
	}

	if param.SuperiorID != nil {
		id := objects.NewIdentifier(*param.SuperiorID)
		superiorID = &id
	}

	return Employee{
		root: entities.NewEmployee(
			id,
			param.Code,
			param.Name,
			param.Password,
			param.ChangelogParam,
		),
		position:    position,
		department:  department,
		superiorID:  superiorID,
		attendances: make([]entities.Attendance, 0),
	}
}

func (employee Employee) InEmployement() bool {
	if employee.position == nil || employee.department == nil {
		return false
	}

	positionID := objects.GetNumberIdentifier(employee.position.ID())
	return positionID > 0
}

func (employee Employee) WorkInDepartment(name string) bool {
	if !employee.InEmployement() {
		return false
	}

	return employee.department.NameEqual(name)
}

// IsRegisterable returned true if id is none,
// which is employee is not exist in database.
func (employee Employee) IsRegisterable() bool {
	return (objects.GetNumberIdentifier(employee.root.ID()) <= 0 ||
		len(employee.root.Code()) <= len(formatPrefixEmployeeCode)) &&
		!employee.isRegistered
}

func (employee *Employee) Register() {
	if !employee.IsRegisterable() {
		return
	}

	employee.isRegistered = true
	employee.addEvent(events.CreateEmployee{
		Employee:   employee.root,
		Position:   employee.position,
		Department: employee.department,
		SuperiorID: employee.superiorID,
	})
}

type ROPosition interface {
	ID() objects.Identifier[int]
	Name() string
	Changelog() objects.Changelog
}

type RODepartment interface {
	ID() objects.Identifier[int]
	Name() string
	Changelog() objects.Changelog
}

type Superior interface {
	ID() objects.Identifier[int]
	Department() RODepartment
}

func (employee *Employee) AssignSuperior(super Superior, newPositionParam PositionParam) error {
	if super.Department() == nil {
		return domainErr.DepartmentNotExist{}
	}

	// this employee will work under the same department as superior.
	superID := super.ID()
	employee.superiorID = &superID

	employee.addEvent(events.UpdateSuperior{
		EmployeeID: employee.root.ID(),
		SuperiorID: superID,
	})

	newPosition := entities.NewPosition(
		// id in param, could be 0, if it's,
		// then create position in this department.
		newPositionParam.ID,
		objects.GetNumberIdentifier(super.Department().ID()),
		newPositionParam.Name,
		newPositionParam.ChangelogParam,
	)

	employee.position = &newPosition
	// attach superior department to subordinate.
	superDep := entities.NewDepartment(
		objects.GetNumberIdentifier(super.Department().ID()),
		super.Department().Name(),
		objects.ChangelogParam{}, // TODO: Not needed yet.
	)
	employee.department = &superDep
	employee.addEvent(events.CreateOrUsePosition{
		DepartmentID: super.Department().ID(),
		PositionName: newPositionParam.Name,
		Changelog:    objects.NewCreateChangelog(employee.root.Code()),
	})

	return nil
}

func (emp Employee) ID() objects.Identifier[int] {
	return emp.root.ID()
}

func (emp Employee) Department() RODepartment {
	return emp.department
}

func (emp Employee) Position() ROPosition {
	return emp.position
}

func (emp *Employee) Delete() error {
	if emp.IsRegisterable() {
		return domainErr.EmployeeNotExist{}
	}

	if emp.isDeleted {
		return domainErr.EmployeeIsDeleted{}
	}

	emp.addEvent(events.DeleteEmployee{
		ID:        emp.ID(),
		Changelog: emp.root.Changelog().UpdatedNow(emp.root.Code()).DeletedNow(),
	})

	emp.isDeleted = true
	return nil
}

func (emp *Employee) ApplyPosition(posParam PositionParam, depParam DepartmentParam) error {
	if emp.IsRegisterable() {
		return domainErr.EmployeeNotExist{}
	}

	if emp.InEmployement() {
		return domainErr.AlreadEmployee{}
	}

	var (
		posNameEmpty  = len(posParam.Name) <= 0
		depParamEmpty = len(depParam.Name) <= 0
	)

	if posNameEmpty && depParamEmpty {
		emp.addEvent(events.ResignEmployee{
			ID:        emp.ID(),
			Changelog: emp.root.Changelog().UpdatedNow(emp.root.Code()),
		})

		return nil
	} else if posNameEmpty || depParamEmpty {
		return domainErr.PositionOrDepartmentNotExist{}
	}

	newPos := entities.NewPosition(0, 0, posParam.Name, posParam.ChangelogParam)
	emp.position = &newPos

	newDep := entities.NewDepartment(0, depParam.Name, depParam.ChangelogParam)
	emp.department = &newDep

	emp.addEvent(events.CreateOrUsePosition{
		DepartmentName: depParam.Name,
		PositionName:   posParam.Name,
		Changelog:      objects.NewCreateChangelog(emp.root.Code()),
	})

	return nil
}

func (emp *Employee) ChangePositionName(name string) error {
	if emp.IsRegisterable() {
		return domainErr.EmployeeNotExist{}
	}

	if !emp.InEmployement() {
		return domainErr.NotAnEmployee{}
	}

	if len(name) <= 0 {
		return domainErr.PositionOrDepartmentNotExist{}
	}

	emp.position.ChangeName(name)
	emp.addEvent(events.UpdatePosition{
		ID:        emp.position.ID(),
		NewName:   name,
		Changelog: emp.position.Changelog().UpdatedNow(emp.root.Code()),
	})

	return nil
}

func (emp *Employee) DeletePosition() error {
	if emp.IsRegisterable() {
		return domainErr.EmployeeNotExist{}
	}

	if !emp.InEmployement() {
		return domainErr.NotAnEmployee{}
	}

	emp.ApplyPosition(PositionParam{}, DepartmentParam{})

	// TODO: Resign all employee who use this position.
	emp.addEvent(events.DeletePosition{
		ID:        emp.position.ID(),
		Changelog: emp.position.Changelog().UpdatedNow(emp.root.Code()),
	})

	return nil
}

func (emp *Employee) ChangeDepartementName(name string) error {
	if emp.IsRegisterable() {
		return domainErr.EmployeeNotExist{}
	}

	if !emp.InEmployement() {
		return domainErr.NotAnEmployee{}
	}

	if len(name) <= 0 {
		return domainErr.PositionOrDepartmentNotExist{}
	}

	emp.department.ChangeName(name)
	emp.addEvent(events.UpdateDepartment{
		ID:        emp.position.ID(),
		NewName:   name,
		Changelog: emp.position.Changelog().UpdatedNow(emp.root.Code()),
	})

	return nil
}

func (emp *Employee) DeleteDepartment() error {
	if emp.IsRegisterable() {
		return domainErr.EmployeeNotExist{}
	}

	if !emp.InEmployement() {
		return domainErr.NotAnEmployee{}
	}

	emp.ApplyPosition(PositionParam{}, DepartmentParam{})

	// TODO: Resign all employee who use this department.
	emp.addEvent(events.DeleteDepartment{
		ID:        emp.department.ID(),
		Changelog: emp.department.Changelog().UpdatedNow(emp.root.Code()),
	})

	return nil
}
