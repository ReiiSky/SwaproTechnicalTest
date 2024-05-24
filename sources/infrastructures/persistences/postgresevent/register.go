package postgresevent

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/events"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
)

func leftPad(str string, targetLength int, padChar rune) string {
	return fmt.Sprintf("%0*s", targetLength, str)
}

func newEmployeeCode() string {
	now := time.Now().UTC()
	year, month, day := now.Date()

	y := leftPad(strconv.Itoa(year-2000), 2, '0')
	m := leftPad(strconv.Itoa(int(month)), 2, '0')
	d := leftPad(strconv.Itoa(int(day)), 2, '0')
	ms := strconv.Itoa(int(now.Nanosecond()))

	return y + m + d + ms
}

type RegisterImpl struct {
	events.CreateEmployee
}

func (impl RegisterImpl) Fn() persistences.EventImplFn {
	return func(ctx context.Context, event domains.IEvent) error {
		var (
			tx     = ctx.Value(persistences.TXContextKey).(*sql.Tx)
			fnSpec = event.(events.CreateEmployee)
			err    error
		)

		var newEmployeeID int
		var employeeCode = newEmployeeCode()

		err = sq.Insert("employee").
			Columns("employee_code", "name", "password").
			Values(employeeCode, fnSpec.Employee.Name(), fnSpec.Employee.Password().Raw()).
			Suffix("RETURNING \"employee_id\"").
			PlaceholderFormat(sq.Dollar).
			RunWith(tx).
			ScanContext(ctx, &newEmployeeID)

		return err
	}
}
