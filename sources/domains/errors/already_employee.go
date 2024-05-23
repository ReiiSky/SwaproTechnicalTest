package errors

type AlreadEmployee struct{}

func (err AlreadEmployee) Error() string {
	return "Already in employment status"
}
