package restrictions

import "time"

type RestrictionType int64

const (
	Delete    RestrictionType = 0
	Ban       RestrictionType = 1
	ShadowBan RestrictionType = 2
)

type RestrictionObject int64

const (
	User    RestrictionObject = 0
	Rating  RestrictionObject = 1
	Post    RestrictionObject = 2
	Comment RestrictionObject = 3
)

type Restriction struct {
	ID            int64
	ValidUntil    *time.Time
	Type          RestrictionType
	ObjectID      int64
	RestrictionOn RestrictionObject
}

func (r *Restriction) fromModel(model restrictionModel) {
	if r == nil {
		return
	}
	*r = Restriction{
		ID:            model.ID,
		ValidUntil:    model.ValidUntil,
		Type:          RestrictionType(model.Type),
		ObjectID:      model.ObjectID,
		RestrictionOn: RestrictionObject(model.RestrictionOn),
	}
}

type GetObjectRestrictions struct {
	ObjectID      *int64
	RestrictionOn RestrictionObject
	Type          *RestrictionType
}
