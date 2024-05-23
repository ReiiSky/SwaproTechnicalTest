package specifications

type GetEmpty struct {
	Name     string
	Password string
}

func (spec GetEmpty) Specname() string {
	return "GetEmpty"
}
