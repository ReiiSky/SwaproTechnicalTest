package postgresevent

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

func upsertChangelog(ctx context.Context, tx *sql.Tx, logname string, id int, log objects.Changelog) {
	if log.UpdatedAt() != nil || log.DeletedAt() != nil {
		updatedField := map[string]interface{}{}

		if log.UpdatedAt() != nil {
			updatedField["updated_at"] = log.UpdatedAt().ToISOUTC()
			updatedField["updated_by"] = *log.UpdatedBy()
		}

		if log.DeletedAt() != nil {
			updatedField["deleted_at"] = log.DeletedAt().ToISOUTC()
		}

		sq.Update("changelog").
			Where(sq.Eq{"log_name": logname}).
			Where(sq.Eq{"id": id}).
			SetMap(updatedField).
			PlaceholderFormat(sq.Dollar).
			RunWith(tx).
			ExecContext(ctx)
	} else {
		sq.Insert("changelog").
			Columns("log_name", "id", "created_at", "created_by").
			Values(logname, id, log.CreatedAt().ToISOUTC(), log.CreatedBy()).
			PlaceholderFormat(sq.Dollar).
			RunWith(tx).
			ExecContext(ctx)
	}
}
