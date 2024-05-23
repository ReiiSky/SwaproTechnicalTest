package getpositioninformation

type PositionInformationOutput struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	DepartmentName string `json:"departement_name"`
	EmployeeCount  int    `json:"employee_count"`
	CreatedAt      string `json:"created_at"`
}
