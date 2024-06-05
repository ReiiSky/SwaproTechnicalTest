package specifications

type GetByID struct {
	ID int

	WithMembership  bool
	AttendanceLimit int
}

func (spec GetByID) Specname() string {
	return "GetByID"
}
