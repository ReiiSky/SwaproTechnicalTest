package entities

import (
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type Attendance struct {
	Entity[int]
	employeeCode objects.InformationNumber[string]
	location     Location
	in           objects.SwaproTime
	out          *objects.SwaproTime
	changelog    objects.Changelog
}

type LocationParam struct {
	ID   int
	Name string
	objects.ChangelogParam
}

func NewAttendance(
	id int,
	employeeCode string,
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
		objects.NewInformationNumber(employeeCode),
		NewLocation(locationParam.ID, locationParam.Name, locationParam.ChangelogParam),
		objects.NewSwaproTime(in),
		outTime,
		objects.NewChangelog(changelog),
	}
}

func (att Attendance) ID() objects.Identifier[int] {
	return att.identifier
}

type ROLocation interface {
	ID() int
	Name() string
	Changelog() objects.Changelog
}

func (att *Attendance) ChangeLocationName(newName string) {
	att.location.ChangeName(att.employeeCode, newName)
}

func (att Attendance) Location() ROLocation {
	return att.location
}

func (att Attendance) In() objects.SwaproTime {
	return att.in
}

func (att *Attendance) CheckOut() {
	now := objects.NewSwaproTimeNow()
	att.out = &now

	att.changelog = att.changelog.UpdatedNow(objects.GetStringInformationNumber(att.employeeCode))
}

func (att Attendance) Out() *objects.SwaproTime {
	return att.out
}

func (att Attendance) Changelog() objects.Changelog {
	return att.changelog
}
