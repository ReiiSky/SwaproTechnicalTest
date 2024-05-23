package specifications

type GetByName struct {
	Name string
}

func (spec GetByName) Specname() string {
	return "GetByName"
}
