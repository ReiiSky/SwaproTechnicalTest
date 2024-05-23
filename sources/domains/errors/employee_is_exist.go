package errors

type EmployeeIsExist struct{}

func (err EmployeeIsExist) Error() string {
	return "Employee is exist"
}
