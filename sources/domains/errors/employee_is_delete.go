package errors

type EmployeeIsDeleted struct{}

func (err EmployeeIsDeleted) Error() string {
	return "Employee is deleted"
}
