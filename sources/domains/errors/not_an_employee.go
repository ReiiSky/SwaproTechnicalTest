package errors

type NotAnEmployee struct{}

func (err NotAnEmployee) Error() string {
	return "Not in employment"
}
