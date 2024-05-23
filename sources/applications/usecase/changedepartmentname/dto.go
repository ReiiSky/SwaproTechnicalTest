package changedepartmentname

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type ChangeDepartmentInput struct {
	NewDepartmentName string `json:"new_department_name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
}

var validatorInstance = validator.New()

func NewChangeDepartmentInput(payload string) (ChangeDepartmentInput, error) {
	input := ChangeDepartmentInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
