package errors

type AttendanceStillActive struct{}

func (err AttendanceStillActive) Error() string {
	return "Attendance still active, not checkouted"
}
