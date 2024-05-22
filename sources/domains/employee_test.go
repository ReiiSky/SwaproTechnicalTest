package domains_test

import (
	"testing"
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/domains"
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

func TestIsEmployee(t *testing.T) {
	var (
		EmployeeID         = 1
		EmployeeSuperiorID = 2
		EmployeeCode       = "22010001"
		EmployeeName       = "test-123"
		EmployeePassword   = "ca6d00e33edff0e9cb3782d31182de33"

		PositionID   = 5
		PositionName = "Junior Backend Developer"

		DepartmentID   = 9
		DepartmentName = "Sales"
	)

	employee := domains.NewEmployee(EmployeeID, domains.EmployeeParam{
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

	if !employee.InEmployement() {
		t.Error("Employee expected to be in employement")
	}

	if !employee.WorkInDepartment(DepartmentName) {
		t.Error("Employee expected to be work in department: " + DepartmentName)
	}
}
