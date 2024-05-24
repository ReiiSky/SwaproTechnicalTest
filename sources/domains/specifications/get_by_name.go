package specifications

type GetByName struct {
	Name string

	AttendanceLimit int
}

func (spec GetByName) Specname() string {
	return "GetByName"
}
