package restrictions

import "time"

type restrictionModel struct {
	ID            int64      `db:"id"`
	ValidUntil    *time.Time `db:"valid_until"`
	Type          int64      `db:"restriction_type"`
	ObjectID      int64      `db:"object_id"`
	RestrictionOn int64      `db:"restriction_on"`
}
