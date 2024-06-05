package model

var (
	EmployeeColumns   = []string{"employee_id", "employee_code", "position_id", "superior_id", "name", "password"}
	PositionColumns   = []string{"position_id", "department_id", "name"}
	DepartmentColumns = []string{"department_id", "department_name"}
	MembershipColumns = []string{"membership_id", "name", "address", "is_active", "employee_id"}
)
