package entities

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type Location struct {
	Entity[int]
	name      string
	changelog objects.Changelog
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
	}
}
