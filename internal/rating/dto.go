package rating

type Rating struct {
	ID        int64
	Score     float32
	Review    string
	FilmID    int64
	UserID    int64
	IsSpecial bool
}

func (r *Rating) SetReview(review string) {
	r.Review = review
}

func (r *Rating) fromModel(model ratingModel) {
	if r == nil {
		return
	}
	*r = Rating{
		ID:        model.ID,
		Score:     model.Score,
		Review:    model.Review,
		FilmID:    model.FilmID,
		UserID:    model.UserID,
		IsSpecial: model.IsSpecial,
	}
}
