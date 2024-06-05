package postgresevent

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/events"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
)

type AddMembershipImpl struct {
	events.AddMembership
}

func (impl AddMembershipImpl) Fn() persistences.EventImplFn {
	return func(ctx context.Context, event domains.IEvent) error {
		var (
			tx      = ctx.Value(persistences.TXContextKey).(*sql.Tx)
			payload = event.(events.AddMembership)
			err     error
		)

		var membershipID int

		err = sq.Insert("membership").
			Columns("employee_id", "name", "password", "address", "is_active").
			Values(
				objects.GetNumberIdentifier(payload.EmployeeID),
				payload.Name,
				payload.Password.Raw(),
				payload.Address,
				payload.IsActive,
			).
			Suffix("RETURNING \"membership_id\"").
			PlaceholderFormat(sq.Dollar).
			RunWith(tx).
			ScanContext(ctx, &membershipID)

		upsertChangelog(ctx, tx, "membership", membershipID, payload.Changelog)

		return err
	}
}
