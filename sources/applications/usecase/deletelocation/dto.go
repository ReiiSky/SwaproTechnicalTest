package deletelocation

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type DeleteLocationInput struct {
	AttendanceID int `json:"attendance_id" validate:"required,gt=0"`
}

var validatorInstance = validator.New()

func NewDeleteLocationInput(payload string) (DeleteLocationInput, error) {
	input := DeleteLocationInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
