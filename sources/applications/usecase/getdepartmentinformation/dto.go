package getdepartmentinformation

type DepartmentInformationOutput struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	EmployeeCount int    `json:"employee_count"`
	PositionCount int    `json:"position_count"`
	CreatedAt     string `json:"created_at"`
}
