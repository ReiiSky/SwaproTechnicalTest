package specifications

type GetByID struct {
	ID int

	AttendanceLimit int
}

func (spec GetByID) Specname() string {
	return "GetByID"
}
