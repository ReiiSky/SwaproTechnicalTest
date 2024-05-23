package register

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type EmployeeRegisterInput struct {
	Name     string `json:"name" validate:"required,minlength=4,maxlength=32"`
	Password string `json:"password" validate:"required,minlength=4,maxlength=32"`
}

var validatorInstance = validator.New()

func NewEmployeeRegisterInput(payload string) (EmployeeRegisterInput, error) {
	input := EmployeeRegisterInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
