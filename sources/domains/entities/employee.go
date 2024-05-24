package entities

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type Employee struct {
	Entity[int]
	code      objects.InformationNumber[string]
	name      string
	password  objects.CryptedInformation
	changelog objects.Changelog
}

func NewEmployee(
	id int,
	code string,
	name,
	password string,
	changelog objects.ChangelogParam,
) Employee {
	return Employee{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		objects.NewInformationNumber(code),
		name,
		objects.NewMD5CryptedInformation(password),
		objects.NewChangelog(changelog),
	}
}

func (employee Employee) Code() string {
	return objects.GetStringInformationNumber(employee.code)
}

func (employee Employee) Name() string {
	return employee.name
}

func (emp Employee) Password() objects.CryptedInformation {
	return emp.password
}

func (emp Employee) Changelog() objects.Changelog {
	return emp.changelog
}
