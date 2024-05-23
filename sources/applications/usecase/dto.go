package usecase

type ErrDTOValidation struct{ msg string }

func NewErrorValidation(err error) ErrDTOValidation {
	return ErrDTOValidation{err.Error()}
}

func (err ErrDTOValidation) Error() string {
	return err.msg
}
