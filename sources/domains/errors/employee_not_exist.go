package errors

type EmployeeNotExist struct{}

func (err EmployeeNotExist) Error() string {
	return "Employee not exist"
}
