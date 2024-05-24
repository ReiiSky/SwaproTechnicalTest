package postgres

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/entities"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
	"github.com/jmoiron/sqlx"
)

var (
	employeeColumns   = []string{"employee_id", "employee_code", "position_id", "superior_id", "name", "password"}
	positionColumns   = []string{"position_id", "department_id", "name"}
	departmentColumns = []string{"department_id", "department_name"}
)

type attendanceRow struct {
	AttendanceID int     `db:"attendance_id"`
	EmployeeID   int     `db:"employee_id"`
	LocationID   int     `db:"location_id"`
	AbsentIn     string  `db:"absent_in"`
	AbsentOut    *string `db:"absent_out"`
	LocationName string  `db:"name"`
}

type departmentRow struct {
	DepartmentID   int    `db:"department_id"`
	DepartmentName string `db:"department_name"`
}

type positionRow struct {
	PositionID   int    `db:"position_id"`
	DepartmentID int    `db:"department_id"`
	Name         string `db:"name"`
}

type employeeRow struct {
	ID         int    `db:"employee_id"`
	Code       string `db:"employee_code"`
	PositionID *int   `db:"position_id"`
	SuperiorID *int   `db:"superior_id"`
	Name       string `db:"name"`
	Password   string `db:"password"`
}

type GetByEmployeeID struct {
	specifications.GetByID
}

func parseStringTimestamp(str *string) *time.Time {
	if str == nil {
		return nil
	}

	t, err := time.Parse(time.RFC3339, *str)

	if err != nil {
		return nil
	}

	return &t
}

func convertChangelogRowToParam(c changelog) objects.ChangelogParam {
	var u *struct {
		At time.Time
		By string
	} = nil

	if c.UpdatedAt != nil {
		u = &struct {
			At time.Time
			By string
		}{
			At: *parseStringTimestamp(c.UpdatedAt),
			By: c.CreatedBy,
		}
	}

	return objects.ChangelogParam{
		CreatedAt: *parseStringTimestamp(&c.CreatedAt),
		CreatedBy: c.CreatedBy,
		Update:    u,
		DeletedAt: parseStringTimestamp(c.DeletedAt),
	}
}

func (impl GetByEmployeeID) Fn() persistences.SpecImplFn {
	return func(ctx context.Context, spec domains.ISpecification) []applications.Aggregate {
		var (
			db     = ctx.Value(persistences.DBContextKey).(*sql.DB)
			fnSpec = spec.(specifications.GetByID)
		)

		rows := []employeeRow{}
		result, _ := sq.
			Select(employeeColumns...).
			From("employee").
			Where(sq.Eq{"employee_id": fnSpec.ID}).
			PlaceholderFormat(sq.Dollar).
			RunWith(db).
			QueryContext(ctx)

		sqlx.StructScan(result, &rows)

		if len(rows) <= 0 {
			return []applications.Aggregate{}
		}

		var (
			positionParam   *domains.PositionParam
			changelogParam  objects.ChangelogParam
			attendanceParam = make([]domains.AttendanceParam, 0)
		)

		employeeData := rows[0]
		changelog := getChangelog(ctx, db, "employee", fnSpec.ID)

		if changelog != nil {
			changelogParam = convertChangelogRowToParam(*changelog)
		}

		if fnSpec.AttendanceLimit > 0 {
			result, _ := sq.Select(
				"attendance.attendance_id",
				"attendance.employee_id",
				"attendance.location_id",
				"attendance.absent_in",
				"attendance.absent_out",
				"location.name",
			).From("attendance").
				InnerJoin("location on attendance.location_id = location.location_id").
				Where(sq.Eq{
					"attendance.employee_id": fnSpec.ID,
				}).
				Limit(uint64(fnSpec.AttendanceLimit)).
				PlaceholderFormat(sq.Dollar).
				RunWith(db).
				QueryContext(ctx)

			atts := []attendanceRow{}
			sqlx.StructScan(result, &atts)

			for _, att := range atts {
				attendanceParam = append(attendanceParam, domains.AttendanceParam{
					ID:        att.AttendanceID,
					AbsentIn:  *parseStringTimestamp(&att.AbsentIn),
					AbsentOut: parseStringTimestamp(att.AbsentOut),
					Location: entities.LocationParam{
						ID:             att.LocationID,
						Name:           att.LocationName,
						ChangelogParam: convertChangelogRowToParam(*getChangelog(ctx, db, "location", att.LocationID)),
					},
					ChangelogParam: convertChangelogRowToParam(*getChangelog(ctx, db, "attendance", att.AttendanceID)),
				})
			}
		}

		if employeeData.PositionID != nil {
			rows := []positionRow{}
			result, _ := sq.
				Select(positionColumns...).
				From("position").
				Where(sq.Eq{"position_id": *employeeData.PositionID}).
				PlaceholderFormat(sq.Dollar).
				RunWith(db).
				QueryContext(ctx)

			sqlx.StructScan(result, &rows)

			if len(rows) > 0 {
				pos := rows[0]
				depRows := []departmentRow{}
				result, _ := sq.
					Select(departmentColumns...).
					From("department").
					Where(sq.Eq{"department_id": pos.DepartmentID}).
					PlaceholderFormat(sq.Dollar).
					RunWith(db).
					QueryContext(ctx)

				sqlx.StructScan(result, &depRows)

				if len(depRows) > 0 {
					dep := depRows[0]
					positionParam = &domains.PositionParam{
						ID:   pos.PositionID,
						Name: pos.Name,
						Department: domains.DepartmentParam{
							ID:             dep.DepartmentID,
							Name:           dep.DepartmentName,
							ChangelogParam: convertChangelogRowToParam(*getChangelog(ctx, db, "department", pos.DepartmentID)),
						},
						EmployeeCount:  1, // not working yet.
						ChangelogParam: convertChangelogRowToParam(*getChangelog(ctx, db, "position", pos.PositionID)),
					}
				}
			}
		}

		employee := domains.NewEmployee(employeeData.ID, domains.EmployeeParam{
			Code:           employeeData.Code,
			Name:           employeeData.Name,
			Password:       employeeData.Password,
			Position:       positionParam,
			SuperiorID:     employeeData.SuperiorID,
			ChangelogParam: changelogParam,
			Attendances:    attendanceParam,
		})

		return []applications.Aggregate{&employee}
	}
}
