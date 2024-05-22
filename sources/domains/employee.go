package domains

import (
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/domains/entities"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
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
