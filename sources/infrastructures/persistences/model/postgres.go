package model

var (
	EmployeeColumns   = []string{"employee_id", "employee_code", "position_id", "superior_id", "name", "password"}
	PositionColumns   = []string{"position_id", "department_id", "name"}
	DepartmentColumns = []string{"department_id", "department_name"}
)
