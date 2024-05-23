package checkin

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type CheckInInput struct {
	LocationName string `json:"location_name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
}

var validatorInstance = validator.New()

func NewCheckInInput(payload string) (CheckInInput, error) {
	input := CheckInInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
