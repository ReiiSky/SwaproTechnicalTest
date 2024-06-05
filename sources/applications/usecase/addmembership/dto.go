package addmembership

import (
	"encoding/json"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/usecase"
	"github.com/go-playground/validator/v10"
)

type AddMembershipInput struct {
	Name    string `json:"name" validate:"required,alphaSpace,minlength=4,maxlength=64"`
	Address string `json:"address" validate:"required,minlength=4,maxlength=64"`
}

var validatorInstance = validator.New()

func NewAddMembershipInput(payload string) (AddMembershipInput, error) {
	input := AddMembershipInput{}
	json.Unmarshal([]byte(payload), &input)

	err := validatorInstance.Struct(validatorInstance)

	if err != nil {
		return input, usecase.NewErrorValidation(err)
	}

	return input, nil
}
