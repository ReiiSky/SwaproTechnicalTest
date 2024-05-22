package entities

import (
	"strings"

	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type Department struct {
	Entity[int]
	name      string
	changelog objects.Changelog
}

func NewDepartment(
	id int,
	name string,
	changelog objects.ChangelogParam,
) Department {
	return Department{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		name,
		objects.NewChangelog(changelog),
	}
}

func (d Department) NameEqual(name string) bool {
	return strings.Compare(d.name, name) == 0
}
