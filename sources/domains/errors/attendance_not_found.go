package errors

type AttendanceNotFound struct{}

func (err AttendanceNotFound) Error() string {
	return "Attendance not found"
}
