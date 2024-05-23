package errors

import "fmt"

type PositionOrDepartmentNotExist struct{ Field string }

func (err PositionOrDepartmentNotExist) Error() string {
	if len(err.Field) <= 0 {
		err.Field = "Name"
	}

	return fmt.Sprintf("Position or Department %s not valid", err.Field)
}
