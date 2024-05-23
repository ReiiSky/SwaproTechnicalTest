package applytoposition

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type ApplyToPositionInput struct {
	PositionName   string `json:"position_name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
	DepartmentName string `json:"department_name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
}

var validatorInstance = validator.New()

func NewApplyToPositionInput(payload string) (ApplyToPositionInput, error) {
	input := ApplyToPositionInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
