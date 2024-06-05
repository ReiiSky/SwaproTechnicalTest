package events

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type AddMembership struct {
	Name       string
	Password   objects.CryptedInformation
	Address    string
	IsActive   bool
	EmployeeID objects.Identifier[int]
	objects.Changelog
}

func (evt AddMembership) Eventname() string {
	return "AddMembership"
}
