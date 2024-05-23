package errors

type DepartmentNotExist struct{}

func (err DepartmentNotExist) Error() string {
	return "Department not exist"
}
