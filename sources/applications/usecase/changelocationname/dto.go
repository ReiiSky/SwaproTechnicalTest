package changelocationname

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type ChangeLocationNameInput struct {
	AttendanceID    int    `json:"attendance_id" validate:"required,gt=0"`
	NewLocationName string `json:"new_location_name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
}

var validatorInstance = validator.New()

func NewChangeLocationNameInput(payload string) (ChangeLocationNameInput, error) {
	input := ChangeLocationNameInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
