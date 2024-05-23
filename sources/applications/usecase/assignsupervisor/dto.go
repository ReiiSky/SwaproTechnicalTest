package assignsupervisor

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type AssignSupervisorInput struct {
	SuperiorID      int    `json:"superior_id" validate:"required,gt=0"`
	NewPositionName string `json:"new_position_name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
}

var validatorInstance = validator.New()

func NewAssignSupervisorInput(payload string) (AssignSupervisorInput, error) {
	input := AssignSupervisorInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
