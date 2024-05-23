package specifications

type GetByID struct {
	ID int
}

func (spec GetByID) Specname() string {
	return "GetByID"
}
