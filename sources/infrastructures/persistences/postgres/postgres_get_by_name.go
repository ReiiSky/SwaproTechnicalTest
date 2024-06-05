package postgres

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/specifications"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences/model"
	"github.com/jmoiron/sqlx"
)

type GetByEmployeeName struct {
	specifications.GetByName
}

func (impl GetByEmployeeName) Fn() persistences.SpecImplFn {
	return func(ctx context.Context, spec domains.ISpecification) []applications.Aggregate {
		var (
			db     = ctx.Value(persistences.DBContextKey).(*sql.DB)
			fnSpec = spec.(specifications.GetByName)
		)

		rows := []employeeRow{}
		result, _ := sq.
			Select(model.EmployeeColumns...).
			From("employee").
			Where(sq.Eq{"name": fnSpec.Name}).
			PlaceholderFormat(sq.Dollar).
			RunWith(db).
			QueryContext(ctx)

		sqlx.StructScan(result, &rows)

		if len(rows) <= 0 {
			return []applications.Aggregate{}
		}

		employee := getEmployeeDetail(ctx, db, rows[0], fnSpec.AttendanceLimit, false)
		return []applications.Aggregate{&employee}
	}
}
