package postgres

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
	"github.com/jmoiron/sqlx"
)

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

		employee := getEmployeeDetail(ctx, db, rows[0], fnSpec.AttendanceLimit)
		return []applications.Aggregate{&employee}
	}
}
