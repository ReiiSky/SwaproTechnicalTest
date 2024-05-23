package usecase

import "errors"

const (
	ErrCodeUnkown = iota
	ErrCodeInvalidRequest
	ErrCodeUnauthorized
	ErrCodeNotFound
	ErrCodeConflict
)

type ErrorWithCode struct {
	ErrCode     int
	ErrInstance error
}

type ErrorMapper struct {
	errors []ErrorWithCode
}

func NewErrorMapper() *ErrorMapper {
	return &ErrorMapper{errors: make([]ErrorWithCode, 0)}
}

func (mapper *ErrorMapper) Add(errCode int, err error) *ErrorMapper {
	mapper.errors = append(mapper.errors, ErrorWithCode{errCode, err})

	return mapper
}

func (mapper ErrorMapper) Map(err error) *ErrorWithCode {
	for _, mappedError := range mapper.errors {
		if errors.As(err, mappedError.ErrInstance) {
			return &mappedError
		}
	}

	return &ErrorWithCode{ErrCodeUnkown, err}
}
