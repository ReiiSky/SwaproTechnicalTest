package getemployeeinfo

type EmployeeRegisterOutput struct {
	Name       string  `json:"name"`
	Position   *string `json:"position_name"`
	Department *string `json:"departement_name"`
	CreatedAt  string  `json:"created_at"`
}
