package entities

import (
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type Attendance struct {
	Entity[int]
	location  Location
	in        objects.SwaproTime
	out       *objects.SwaproTime
	changelog objects.Changelog
}

type LocationParam struct {
	ID   int
	Name string
	objects.ChangelogParam
}

func NewAttendance(
	id int,
	locationParam LocationParam,
	in time.Time,
	out *time.Time,
	changelog objects.ChangelogParam,
) Attendance {
	var outTime *objects.SwaproTime

	if out != nil {
		o := objects.NewSwaproTime(*out)
		outTime = &o
	}

	return Attendance{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		NewLocation(locationParam.ID, locationParam.Name, locationParam.ChangelogParam),
		objects.NewSwaproTime(in),
		outTime,
		objects.NewChangelog(changelog),
	}
}

func (att Attendance) ID() int {
	return objects.GetNumberIdentifier(att.identifier)
}

type ROLocation interface {
	ID() int
	Name() string
}

func (att Attendance) Location() ROLocation {
	return att.location
}

func (att Attendance) In() objects.SwaproTime {
	return att.in
}

func (att Attendance) Out() *objects.SwaproTime {
	return att.out
}

func (att Attendance) Changelog() objects.Changelog {
	return att.changelog
}
