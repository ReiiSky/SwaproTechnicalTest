package getemployeeinfo

type AttendanceOutput struct {
	ID  int     `json:"attendance_id"`
	In  string  `json:"in"`
	Out *string `json:"out"`
}

type EmployeeInfoOutput struct {
	Name       string             `json:"name"`
	Position   *string            `json:"position_name"`
	Department *string            `json:"departement_name"`
	Attendaces []AttendanceOutput `json:"attendances"`
	CreatedAt  string             `json:"created_at"`
}
