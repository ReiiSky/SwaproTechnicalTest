package entities

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type Attendance struct {
	Entity[int]
	employeeID objects.Identifier[int]
	locationID objects.Identifier[int]
	name       string
	in         objects.SwaproTime
	out        *objects.SwaproTime
	changelog  objects.Changelog
}
