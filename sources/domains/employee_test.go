package domains_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/events"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

func newChangelogParam(createdBy string, withDeletedAt, withUpdatedAt bool) objects.ChangelogParam {
	now := time.Now()
	log := objects.ChangelogParam{
		CreatedAt: now,
		CreatedBy: createdBy,
	}

	if withDeletedAt {
		log.DeletedAt = &now
	}

	if withUpdatedAt {
		log.Update = &struct {
			At time.Time
			By string
		}{
			At: now,
			By: createdBy,
		}
	}

	return log
}

var (
	EmployeeID         = 1
	EmployeeSuperiorID = 2
	EmployeeCode       = "22010001"
	EmployeeName       = "test 123"
	EmployeePassword   = "ca6d00e33edff0e9cb3782d31182de33"

	PositionID   = 5
	PositionName = "Junior Backend Developer"

	DepartmentID   = 9
	DepartmentName = "Sales"

	firstEmployee = domains.NewEmployee(EmployeeID, domains.EmployeeParam{
		Code:           EmployeeCode,
		Name:           EmployeeName,
		Password:       EmployeePassword,
		SuperiorID:     &EmployeeSuperiorID,
		ChangelogParam: newChangelogParam(EmployeeCode, true, true),
		Position: &domains.PositionParam{
			ID:             PositionID,
			Name:           PositionName,
			ChangelogParam: newChangelogParam(EmployeeCode, true, true),
			Department: domains.DepartmentParam{
				ID:             DepartmentID,
				Name:           DepartmentName,
				ChangelogParam: newChangelogParam(EmployeeCode, true, true),
			},
		},
		Attendances: []domains.AttendanceParam{},
	})
)

func TestIsEmployee(t *testing.T) {
	if !firstEmployee.InEmployement() {
		t.Error("Employee expected to be in employement")
	}

	if !firstEmployee.WorkInDepartment(DepartmentName) {
		t.Error("Employee expected to be work in department: " + DepartmentName)
	}
}

func TestRegisterUser(t *testing.T) {
	unregisteredEmployee := domains.NewEmployee(0, domains.EmployeeParam{
		Code:     "",
		Name:     "employee 2",
		Password: "af74a83ae0d5777401f86af4df941e98",
	})

	employees := []domains.Employee{firstEmployee, unregisteredEmployee}
	isRegisterableEmployees := []bool{false, true}

	for idx, employee := range employees {
		if employee.IsRegisterable() != isRegisterableEmployees[idx] {
			t.Error("Employee expected not to be registerble, index: " + strconv.Itoa(idx))

			return
		}

		if !employee.IsRegisterable() {
			continue
		}

		employee.Register()
		employeeEvents := employee.Events()

		if len(employeeEvents) != 1 {
			t.Error("Employee aggregate supposed to have one event")
		}

		if _, ok := employeeEvents[0].Top().(events.CreateEmployee); !ok {
			t.Error("Aggregate inside employee after register is not CreateEmployee, expected CreateEmployee.")
		}
	}
}

func TestAssignSupervisor(t *testing.T) {
	employeeIDs := []int{7, 9}
	employeeCode := []string{"22010007", "22010009"}

	employeeParams := make([]domains.EmployeeParam, 2)
	employeeParams[0] = domains.EmployeeParam{
		Code:     employeeCode[0],
		Name:     "employee seven",
		Password: "7eafc9a76b469dcbd0a6b3b4b79870da",
		Position: &domains.PositionParam{
			ID:   2,
			Name: "Tech Lead",
			Department: domains.DepartmentParam{
				ID:   2,
				Name: "IT",
			},
			ChangelogParam: newChangelogParam(employeeCode[0], false, false),
		},
		ChangelogParam: newChangelogParam(employeeCode[0], false, false),
	}

	employeeParams[1] = domains.EmployeeParam{
		Code:     employeeCode[1],
		Name:     "employee nine",
		Password: "500e869096404a205a2b186a36fe0867",
		// the above employee wil be a supervisor of this employee.
		ChangelogParam: newChangelogParam(employeeCode[0], false, false),
	}

	employees := make([]domains.Employee, len(employeeParams))

	for idx, param := range employeeParams {
		employees[idx] = domains.NewEmployee(employeeIDs[idx], param)
	}

	superior := employees[0]
	subordinate := employees[1]

	err := subordinate.AssignSuperior(superior, domains.PositionParam{
		Name: "Data Scientist",
		Department: domains.DepartmentParam{
			Name: "Growth",
		},
	})

	if err != nil {
		t.Error("Assign superior error is not a nill with message: " + err.Error())
	}

	subEvents := subordinate.Events()

	// events of assign superior is update superior and
	// update position.
	if len(subEvents) != 2 {
		t.Error("Employee aggregate supposed to have two event")
	}

	if _, ok := subEvents[0].Top().(events.UpdateSuperior); !ok {
		t.Error("Aggregate inside employee after assign superior is not UpdateSuperior.")
	}

	if _, ok := subEvents[1].Top().(events.CreateOrUsePosition); !ok {
		t.Error("Aggregate inside employee after assign superior is not CreateOrUsePosition.")
	}
}
