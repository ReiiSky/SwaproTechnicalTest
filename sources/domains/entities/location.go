package entities

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type Location struct {
	Entity[int]
	name      string
	changelog objects.Changelog

	// internal state
	isDeleted bool
}

func NewLocation(
	id int,
	name string,
	changelog objects.ChangelogParam,
) Location {
	return Location{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		name,
		objects.NewChangelog(changelog),
		false,
	}
}

func (loc Location) ID() objects.Identifier[int] {
	return loc.identifier
}

func (loc *Location) ChangeName(employeeCode objects.InformationNumber[string], newName string) {
	loc.name = newName
	loc.changelog = loc.changelog.
		UpdatedNow(objects.GetStringInformationNumber(employeeCode))
}

func (loc *Location) Delete(employeeCode objects.InformationNumber[string]) {
	if loc.isDeleted {
		return
	}

	loc.isDeleted = true
	loc.changelog = loc.changelog.
		UpdatedNow(objects.GetStringInformationNumber(employeeCode)).
		DeletedNow()
}

func (loc Location) Name() string {
	return loc.name
}

func (loc Location) Changelog() objects.Changelog {
	return loc.changelog
}
