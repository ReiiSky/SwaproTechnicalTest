package postgresevent

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

func upsertChangelog(ctx context.Context, tx *sql.Tx, logname string, id int, log objects.Changelog) {
	var (
		updatedAt *string
		updatedBy *string
		deletedAt *string
	)

	if log.UpdatedAt() != nil {
		u := log.UpdatedAt().ToISOUTC()
		updatedAt = &u
	}

	if log.UpdatedBy() != nil {
		updatedBy = log.UpdatedBy()
	}

	if log.DeletedAt() != nil {
		da := log.DeletedAt().ToISOUTC()
		deletedAt = &da
	}

	sq.Insert("changelog").
		Columns("log_name", "id", "created_at", "created_by").
		Values(logname, id, log.CreatedAt(), log.CreatedBy()).
		Suffix(fmt.Sprintf(
			"on conflict (log_name, id) do update set updated_at = '%v',"+
				"updated_by = '%v', deleted_at = '%v'",
			updatedAt, updatedBy, deletedAt),
		).
		PlaceholderFormat(sq.Dollar).
		RunWith(tx).
		ExecContext(ctx)
}
